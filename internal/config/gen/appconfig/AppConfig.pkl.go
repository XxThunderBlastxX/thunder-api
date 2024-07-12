// Code generated from Pkl module `thunderblast.thunderapi.AppConfig`. DO NOT EDIT.
package appconfig

import (
	"context"

	"github.com/XxThunderBlastxX/thunder-api/internal/config/gen/cloudflareconfig"
	"github.com/XxThunderBlastxX/thunder-api/internal/config/gen/databaseconfig"
	"github.com/apple/pkl-go/pkl"
)

type AppConfig struct {
	Name string `pkl:"name"`

	Version string `pkl:"version"`

	Port string `pkl:"port"`

	Db *databaseconfig.DatabaseConfig `pkl:"db"`

	Cloudflare *cloudflareconfig.CloudflareConfig `pkl:"cloudflare"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a AppConfig
func LoadFromPath(ctx context.Context, path string) (ret *AppConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a AppConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*AppConfig, error) {
	var ret AppConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
