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
func (b *backend) tokenRenew(context context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {

	id := d.Get("id").(string)

	entity, err := req.Storage.Get(context, id+"/auth_token")
	if err != nil {
		return nil, err
	}
	authToken := string(entity.Value)

	bldrEntity, err := req.Storage.Get(context, id+"/bldr_url")
	if err != nil {
		return nil, err
	}
	habBldrUrl := string(bldrEntity.Value)
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
	request.Header.Set("Authorization", "Bearer "+authToken)
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	b.Logger().Info("Returned response code " + string(resp.StatusCode))

	var result map[string]string

	json.NewDecoder(resp.Body).Decode(&result)
	token := result["token"]

	tokenStorage := logical.StorageEntry{
		Key:      id + "/auth_token",
		Value:    []byte(token),
		SealWrap: false,
	}

	req.Storage.Put(context, &tokenStorage)

	urlStorage := logical.StorageEntry{
		Key:      id + "/bldr_url",
		Value:    []byte(habBldrUrl),
		SealWrap: false,
	}

	req.Storage.Put(context, &urlStorage)

	return &logical.Response{
		Data: map[string]interface{}{
			"id":         id,
			"auth_token": token,
			"bldr_url":   habBldrUrl,
		},
	}, nil
}
