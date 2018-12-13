package habtoken

import (
	"context"
	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/framework"
)

type backend struct {
	*framework.Backend
}

func Factory(ctx context.Context, conf *logical.BackendConfig) (logical.Backend, error) {
	b, err := Backend(conf)
	if err != nil {
		return nil, err
	}
	if err := b.Setup(ctx, conf); err != nil {
		return nil, err
	}
	return b, nil
}

func Backend(conf *logical.BackendConfig) (*backend, error) {
	var b backend
	b.Backend = &framework.Backend{
		BackendType: logical.TypeLogical,
		Help: `
The gen secrets engine generates passwords and passphrases, and optionally
stores the resulting password in an accessor.
		`,
		Paths: []*framework.Path{
			// habtoken/rotate
			&framework.Path{
				Pattern:      "rotate",
				HelpSynopsis: "Returns the rotated token",
				HelpDescription: `
rotates the auth token
`,
				Fields: map[string]*framework.FieldSchema{
					"current": &framework.FieldSchema{
						Type:        framework.TypeString,
						Description: "The Current Token",
					},
				},
				Callbacks: map[logical.Operation]framework.OperationFunc{
					logical.UpdateOperation: b.tokenRotate,
				},
			},
		},
	}

	return &b
}

const backendHelp = `
Interacts with the habitat depot to rotate depot tokens.
`
