package habtoken

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/framework"
	// "github.com/pkg/errors"
)

// pathPassword corresponds to POST gen/password.
func (b *backend) tokenRotate(context context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {
	// if err := validateFields(req, d); err != nil {
	// 	return nil, logical.CodedError(http.StatusUnprocessableEntity, err.Error())
	// }

	b.Logger().Info("Hab Depot Logger")
	current := d.Get("current").(string)
	if current == "" {
		entity, err := req.Storage.Get(context, "token")
		if err != nil {
			return nil, err
		}
		current = string(entity.Value)
		b.Logger().Info("Checking storage backend " + current)
	}

	habBldrUrl := d.Get("hab_bldr_url").(string)
	apiPath := "/v1/profile/access-tokens"

	b.Logger().Info("Making request against: " + habBldrUrl)

	// skip tls verify
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	request, err := http.NewRequest("POST", habBldrUrl+apiPath, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+current)
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	b.Logger().Info("Returned response code " + string(resp.StatusCode))

	var result map[string]string

	json.NewDecoder(resp.Body).Decode(&result)
	token := result["token"]

	item := logical.StorageEntry{
		Key:      "token",
		Value:    []byte(token),
		SealWrap: false,
	}

	req.Storage.Put(context, &item)

	return &logical.Response{
		Data: map[string]interface{}{
			"value": token,
		},
	}, nil
}
