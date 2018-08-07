package secretsejson

import (
	"context"

	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/framework"
)

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
		Paths: framework.PathAppend(
			ejsonKeysPaths(&b),
			ejsonPaths(&b),
		),
		Secrets:     []*framework.Secret{},
		BackendType: logical.TypeLogical,
	}
	return &b
}

type backend struct {
	*framework.Backend
}