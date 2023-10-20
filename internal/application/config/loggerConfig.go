package config

import (
	"path"
	"sync"
)

const loggerConfigFilename = "logger.config.yaml"

type LoggerConfig struct {
	InfoLogPath  string `yaml:"infoLogPath" env-required:"true"`
	ErrorLogPath string `yaml:"errorLogPath" env-required:"true"`
	TraceLogPath string `yaml:"traceLogPath" env-required:"true"`
	FatalLogPath string `yaml:"fatalLogPath" env-required:"true"`

	InfoLogAbsPath, ErrorLogAbsPath, TraceLogAbsPath, FatalLogAbsPath string
}

var (
	loggerCfgInst     = &LoggerConfig{}
	loadLoggerCfgOnce = &sync.Once{}
)

func Logger() LoggerConfig {
	loadLoggerCfgOnce.Do(func() {
		readConfig(path.Join(Env().ConfigAbsPath, loggerConfigFilename), loggerCfgInst)
		loggerCfgInst.InfoLogAbsPath = path.Join(Env().ProjectAbsPath, loggerCfgInst.InfoLogPath)
		loggerCfgInst.ErrorLogAbsPath = path.Join(Env().ProjectAbsPath, loggerCfgInst.ErrorLogPath)
		loggerCfgInst.TraceLogPath = path.Join(Env().ProjectAbsPath, loggerCfgInst.TraceLogPath)
		loggerCfgInst.FatalLogPath = path.Join(Env().ProjectAbsPath, loggerCfgInst.FatalLogAbsPath)
	})

	return *loggerCfgInst
}
