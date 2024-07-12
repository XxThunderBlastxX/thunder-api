// Code generated from Pkl module `thunderblast.thunderapi.CloudflareConfig`. DO NOT EDIT.
package cloudflareconfig

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type CloudflareConfig struct {
	AccountID string `pkl:"accountID"`

	KvNamespaceID string `pkl:"kvNamespaceID"`

	Token string `pkl:"token"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a CloudflareConfig
func LoadFromPath(ctx context.Context, path string) (ret *CloudflareConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a CloudflareConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*CloudflareConfig, error) {
	var ret CloudflareConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
