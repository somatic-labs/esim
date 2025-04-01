package cmd

import (
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
)

type AppConfig struct {
	serverconfig.Config
}

func NewDefaultConfig() *AppConfig {
	return &AppConfig{
		Config: *serverconfig.DefaultConfig(),
	}
}

func DefaultConfigTemplate() string {
	return serverconfig.DefaultConfigTemplate
}
