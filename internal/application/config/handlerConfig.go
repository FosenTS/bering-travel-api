package config

import (
	"path"
	"sync"
)

const handlerConfigFilename = "handler.config.yaml"

type HandlerConfig struct {
	DefaultTableCapacity int    `yaml:"defaultTableCapacity" env-required:"true"` //env-default:"16"`
	DefaultSortOrder     string `yaml:"defaultSortOrder" env-required:"true"`     //ASC'
}

var (
	handlerCfgInst     = &HandlerConfig{}
	loadHandlerCfgOnce = sync.Once{}
)

func Handler() HandlerConfig {
	loadHandlerCfgOnce.Do(func() {
		readConfig(path.Join(Env().ConfigAbsPath, handlerConfigFilename), handlerCfgInst)
	})

	return *handlerCfgInst
}
