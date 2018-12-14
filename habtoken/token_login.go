package habtoken

import (
	"context"

	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/framework"
)

func (b *backend) tokenLogin(context context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {

	id := d.Get("id").(string)
	authToken := d.Get("auth_token").(string)
	habBldrUrl := d.Get("hab_bldr_url").(string)

	tokenStorage := logical.StorageEntry{
		Key:      id + "/auth_token",
		Value:    []byte(authToken),
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
			"auth_token": authToken,
			"bldr_url":   habBldrUrl,
		},
	}, nil
}
