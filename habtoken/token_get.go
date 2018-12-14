package habtoken

import (
	"context"

	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/framework"
)

func (b *backend) tokenGet(context context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {

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

	return &logical.Response{
		Data: map[string]interface{}{
			"id":         id,
			"auth_token": authToken,
			"bldr_url":   habBldrUrl,
		},
	}, nil
}
