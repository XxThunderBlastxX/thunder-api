package env

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// Env is a struct that holds the environment variables
// for the application
type Env struct {
	APIPort       string `mapstructure:"PORT"`
	CFToken       string `mapstructure:"CF_TOKEN"`
	CFAccountId   string `mapstructure:"CF_ID"`
	KvNamespaceId string `mapstructure:"KV_NAMESPACE_ID"`
}

// LoadEnv loads the environment variables from the given path
func LoadEnv(path string) (env Env, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return env, err
	}

	if err := viper.Unmarshal(&env, func(config *mapstructure.DecoderConfig) {
		config.ErrorUnused = true
		config.ErrorUnset = true
	}); err != nil {
		return env, err
	}

	return
}
