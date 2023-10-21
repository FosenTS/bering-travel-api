package config

import (
	"github.com/achillescres/pkg/gconfig"
	"github.com/sirupsen/logrus"
)

func readConfig(cfgPath string, cfgInst gconfig.Config) {
	logrus.Infof("reading %s\n", cfgPath)
	err := gconfig.ReadConfig(cfgPath, cfgInst)
	if err != nil {
		logrus.Fatalf("fatal reading configs with path %s: %s\n", cfgPath, err)
	}
	logrus.Infof("successfully read %s\n", cfgPath)
}
