package config

import (
	"github.com/achillescres/pkg/gconfig"
	"github.com/achillescres/pkg/mysync"
	log "github.com/sirupsen/logrus"
	"path"
)

const (
	envLocalFilename = ".env.local"
	envDevFilename   = ".env.dev"
	envProdFilename  = ".env.prod"
)

// EnvConfig Env
type EnvConfig struct {
	ProjectAbsPath string `env:"PROJECT_ABS_PATH" env-required:"true"`
	ConfigPath     string `env:"CONFIG_PATH" env-required:"true"`

	// HTTP
	HttpIP   string `env:"HTTP_IP" env-required:"true"`
	HttpPort string `env:"HTTP_PORT" env-required:"true"`

	// Postgres
	PostgresHost         string `env:"POSTGRES_HOST" env-required:"true"`
	PostgresPort         string `env:"POSTGRES_PORT" env-required:"true"`
	PostgresUsername     string `env:"POSTGRES_USERNAME" env-required:"true"`
	PostgresPassword     string `env:"POSTGRES_PASSWORD" env-required:"true"`
	PostgresDatabaseName string `env:"POSTGRES_DATABASE_NAME" env-required:"true"`
	PostgresUseCA        bool   `env:"POSTGRES_USE_CA" env-required:"true"`
	PostgresCaPath       string `env:"POSTGRES_CA_PATH" env_required:"true"`

	// FTP tais output
	FTPTaisOutputHost     string `env:"FTP_TAIS_OUTPUT_HOST" env_required:"true"`
	FTPTaisOutputPort     uint   `env:"FTP_TAIS_OUTPUT_PORT" env-required:"true"`
	FTPTaisOutputUsername string `env:"FTP_TAIS_OUTPUT_USERNAME" env_required:"true"`
	FTPTaisOutputPassword string `env:"FTP_TAIS_OUTPUT_PASSWORD" env_required:"true"`

	// Artem forecast
	ArtemForecastContainerURL string `env:"ARTEM_FORECAST_CONTAINER_URL" env-required:"true"`

	//External API
	//reports
	ReportsContainerURL string `env:"REPORTS_CONTAINER_URL" env-required:"true"`

	// Secrets
	PasswordHashSalt string `env:"PASSWORD_HASH_SALT" env-required:"true"`
	JWTSecret        string `env:"JWT_SECRET" env-required:"true"`

	ConfigAbsPath string
}

var (
	envCfgInst  = &EnvConfig{}
	loadEnvOnce = mysync.NewOnce()
)

func LoadEnv(mode string) {
	envFilename := ""
	switch mode {
	case Mode.Local():
		envFilename = envLocalFilename
	case Mode.Dev():
		envFilename = envDevFilename
	case Mode.Prod():
		envFilename = envProdFilename
	default:
		panic("invalid mode")
	}
	loadEnvOnce.Do(func() {
		err := gconfig.ReadEnv(envFilename, envCfgInst)
		if err != nil {
			log.Fatalf("fatal reading env: %s\n", err)
		}

		envCfgInst.ConfigAbsPath = path.Join(envCfgInst.ProjectAbsPath, envCfgInst.ConfigPath)

		log.Infoln("Env successfully read")
	})
}

func Env() EnvConfig {
	if !loadEnvOnce.IsDone() {
		log.Fatalln("env must be loaded before use!")
	}
	return *envCfgInst
}
