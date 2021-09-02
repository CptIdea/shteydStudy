package config

import "github.com/kelseyhightower/envconfig"

func GetRefConfig() (RefConfig, error) {
	var cfg RefConfig

	err := envconfig.Process("reference", &cfg)
	if err != nil {
		return RefConfig{}, err
	}

	return cfg, nil
}
