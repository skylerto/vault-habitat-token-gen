package habtoken

import (
	"context"
	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/framework"
)

// New returns a new backend as an interface. This func
// is only necessary for builtin backend plugins.
func New() (interface{}, error) {
	return Backend(), nil
}

// Factory returns a new backend as logical.Backend.
func Factory(ctx context.Context, conf *logical.BackendConfig) (logical.Backend, error) {
	b := Backend()
	if err := b.Setup(ctx, conf); err != nil {
		return nil, err
	}
	return b, nil
}

// FactoryType is a wrapper func that allows the Factory func to specify
// the backend type for the mock backend plugin instance.
func FactoryType(backendType logical.BackendType) logical.Factory {
	return func(ctx context.Context, conf *logical.BackendConfig) (logical.Backend, error) {
		b := Backend()
		b.BackendType = backendType
		if err := b.Setup(ctx, conf); err != nil {
			return nil, err
		}
		return b, nil
	}
}

// Backend returns a private embedded struct of framework.Backend.
func Backend() *backend {
	var b backend
	b.Backend = &framework.Backend{
		Help: "",
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
					"hab_bldr_url": &framework.FieldSchema{
						Type:        framework.TypeString,
						Description: "The Habitat Builder to Auth against",
					},
				},
				Callbacks: map[logical.Operation]framework.OperationFunc{
					logical.UpdateOperation: b.tokenRotate,
				},
			},
		},
		PathsSpecial: &logical.Paths{
			Unauthenticated: []string{
				"special",
			},
		},
		Secrets:     []*framework.Secret{},
		Invalidate:  b.invalidate,
		BackendType: logical.TypeLogical,
	}
	b.internal = "bar"
	return &b
}

type backend struct {
	*framework.Backend

	// internal is used to test invalidate
	internal string
}

func (b *backend) invalidate(ctx context.Context, key string) {
	switch key {
	case "internal":
		b.internal = ""
	}
}

// func Factory(ctx context.Context, conf *logical.BackendConfig) (logical.Backend, error) {
// 	b, err := Backend(conf)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err := b.Setup(ctx, conf); err != nil {
// 		return nil, err
// 	}
// 	return b, nil
// }

// func Backend(conf *logical.BackendConfig) (*backend, error) {
// 	var b backend
// 	b.Backend = &framework.Backend{
// 		BackendType: logical.TypeLogical,
// 		Help: `
// The gen secrets engine generates passwords and passphrases, and optionally
// stores the resulting password in an accessor.
// 		`,
// 		Paths: []*framework.Path{
// 			// habtoken/rotate
// 			&framework.Path{
// 				Pattern:      "rotate",
// 				HelpSynopsis: "Returns the rotated token",
// 				HelpDescription: `
// rotates the auth token
// `,
// 				Fields: map[string]*framework.FieldSchema{
// 					"current": &framework.FieldSchema{
// 						Type:        framework.TypeString,
// 						Description: "The Current Token",
// 					},
// 				},
// 				Callbacks: map[logical.Operation]framework.OperationFunc{
// 					logical.UpdateOperation: b.tokenRotate,
// 				},
// 			},
// 		},
// 	}

// 	return &b, nil
// }

const backendHelp = `
Interacts with the habitat depot to rotate depot tokens.
`
