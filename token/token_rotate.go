package habtoken

import (
	"context"
	"net/http"

	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/framework"
	"github.com/pkg/errors"
)

// pathPassword corresponds to POST gen/password.
func (b *backend) tokenRotate(context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {
	if err := validateFields(req, d); err != nil {
		return nil, logical.CodedError(http.StatusUnprocessableEntity, err.Error())
	}

	current := d.Get("current").(string)
	token := "this is a token"

	return &logical.Response{
		Data: map[string]interface{}{
			"value": token,
		},
	}, nil
}
