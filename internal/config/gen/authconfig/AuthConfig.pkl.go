// Code generated from Pkl module `thunderblast.thunderapi.AuthConfig`. DO NOT EDIT.
package authconfig

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type AuthConfig struct {
	Domain string `pkl:"domain"`

	Audience string `pkl:"audience"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a AuthConfig
func LoadFromPath(ctx context.Context, path string) (ret *AuthConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a AuthConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*AuthConfig, error) {
	var ret AuthConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
