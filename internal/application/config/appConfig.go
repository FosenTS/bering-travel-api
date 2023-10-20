package config

import (
	"path"
	"sync"
	"time"
)

const appConfigFilename = "app.config.yaml"

type AppConfig struct {
	IsDev                  bool `yaml:"isDev" env-required:"true"`
	TryRevoke              bool `yaml:"tryRevoke" env-required:"true"`
	RevokeStunMilliseconds uint `yaml:"revokeStunMilliseconds" env-required:"true"`

	HTTP struct {
		IP   string
		Port string
	}
	RevokeStunDuration time.Duration
}

var (
	appCfgInst     = &AppConfig{}
	loadAppCfgOnce = sync.Once{}
)

func App() AppConfig {
	loadAppCfgOnce.Do(func() {
		readConfig(path.Join(Env().ConfigAbsPath, appConfigFilename), appCfgInst)
		appCfgInst.RevokeStunDuration = time.Duration(appCfgInst.RevokeStunMilliseconds) * time.Millisecond
		appCfgInst.HTTP.IP = Env().HttpIP
		appCfgInst.HTTP.Port = Env().HttpPort
	})

	return *appCfgInst
}
