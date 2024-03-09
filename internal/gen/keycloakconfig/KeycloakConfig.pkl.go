// Code generated from Pkl module `thunderblast.thunderapi.KeycloakConfig`. DO NOT EDIT.
package keycloakconfig

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type KeycloakConfig struct {
	Realm string `pkl:"realm"`

	AuthUrl string `pkl:"authUrl"`

	ClientId string `pkl:"clientId"`

	ClientSecret string `pkl:"clientSecret"`

	RealmRSA256PublicKey string `pkl:"realmRSA256PublicKey"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a KeycloakConfig
func LoadFromPath(ctx context.Context, path string) (ret *KeycloakConfig, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a KeycloakConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*KeycloakConfig, error) {
	var ret KeycloakConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
