package habtoken

import (
	"context"
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

	current := d.Get("current").(string)
	habBldrUrl := d.Get("hab_bldr_url").(string)
	apiPath := "/v1/profile/access-tokens"

	client := &http.Client{}
	request, _ := http.NewRequest("POST", habBldrUrl+apiPath, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+current)
	resp, _ := client.Do(request)

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	token := result["token"]

	return &logical.Response{
		Data: map[string]interface{}{
			"value": token,
		},
	}, nil
}
