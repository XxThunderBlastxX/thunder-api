// Code generated from Pkl module `thunderblast.thunderapi.SMTPConfig`. DO NOT EDIT.
package smtpconfig

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type SMTPConfig struct {
	Host string `pkl:"host"`

	Port string `pkl:"port"`

	User string `pkl:"user"`

	Password string `pkl:"password"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a SMTPConfig
func LoadFromPath(ctx context.Context, path string) (ret *SMTPConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a SMTPConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*SMTPConfig, error) {
	var ret SMTPConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
