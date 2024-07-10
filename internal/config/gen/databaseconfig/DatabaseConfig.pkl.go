// Code generated from Pkl module `thunderblast.thunderapi.DatabaseConfig`. DO NOT EDIT.
package databaseconfig

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type DatabaseConfig struct {
	Dsn string `pkl:"dsn"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a DatabaseConfig
func LoadFromPath(ctx context.Context, path string) (ret *DatabaseConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a DatabaseConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*DatabaseConfig, error) {
	var ret DatabaseConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
