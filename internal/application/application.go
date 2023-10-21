package application

import (
	"bering-travel-api/internal/application/config"
	"bering-travel-api/internal/application/product"
	"context"
	"fmt"
	"github.com/achillescres/pkg/db/postgresql"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"os"
	"time"
)

type App interface {
	Run(ctx context.Context) error
	runHTTP() error
}

type app struct {
	cfg       config.AppConfig
	pgPool    postgresql.PGXPool
	endpoint  *product.Endpoint
	logEntry  *log.Entry
	ginLogger gin.HandlerFunc
}

func NewApp(ctx context.Context, logEntry *log.Entry, ginLogger gin.HandlerFunc) (App, error) {

	// Build postgres pool
	log.Infoln("Creating pgxpool")
	postgresConfig := config.Postgres()
	pgPool, err := postgresql.NewPGXPool(
		ctx,
		&postgresql.ClientConfig{
			MaxConnections:        postgresConfig.MaxConnections,
			MaxConnectionAttempts: postgresConfig.MaxConnectionAttempts,
			WaitingDuration:       postgresConfig.WaitTimeout,
			Username:              postgresConfig.Username,
			Password:              postgresConfig.Password,
			Host:                  postgresConfig.Host,
			Port:                  postgresConfig.Port,
			DatabaseName:          postgresConfig.DatabaseName,
			UseCA:                 postgresConfig.UseCA,
			CaAbsPath:             postgresConfig.CAAbsPath,
			SimpleQueryMode:       postgresConfig.SimpleQueryMode,
		},
		logEntry.WithField("location", "pgxpool"),
	)
	if err != nil {
		log.Errorf("error creating pgxpool: %s\n", err)
		return nil, err
	}
	log.Warnf("This is pgxpool: %v\n", pgPool)

	// Create repositories
	log.Infoln("Creating repository...")
	repos := product.NewStorages(pgPool)

	// Create services
	log.Infoln("Creating services...")
	services, err := product.NewServices(repos)
	if err != nil {
		return nil, fmt.Errorf("fatal couldn't create services: %w", err)
	}

	// Create gateways
	log.Infoln("Gateways layer")
	gateways, err := product.NewGateways(
		services,
	)

	// Controllers
	log.Infoln("Creating controllers")
	controllers, err := product.NewControllers(
		gateways,
		config.Handler(),
	)
	if err != nil {
		return nil, fmt.Errorf("fatal couldn't create controllers: %s", err)
	}

	// Endpoint artificial-layer
	log.Infoln("Creating artificial endpoint layer")
	endpoint := product.NewEndpoint(controllers)

	return &app{
		cfg:       config.App(),
		endpoint:  endpoint,
		pgPool:    pgPool,
		logEntry:  logEntry,
		ginLogger: ginLogger,
	}, nil
}

func (app *app) Run(ctx context.Context) error {

	fmt.Println("AA")

	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(func() error {
		return app.runHTTP()
	})

	return grp.Wait()
}

func (app *app) runHTTP() error {
	log.Infoln("Start HTTP")
	r := gin.New()
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: os.Stdout}), gin.Recovery())

	err := app.endpoint.RegisterHandlersToGroup(r)
	if err != nil {
		return err
	}

	addr := app.cfg.HTTP.IP
	if len(app.cfg.HTTP.Port) != 0 {
		addr += ":" + app.cfg.HTTP.Port
	}

	for {
		err = r.Run(addr)
		if err != nil {
			log.Errorf("error running http server: %s", err)
		} else {
			log.Warnln("warn running http server: it stopped without error")
		}
		if app.cfg.IsDev {
			log.Errorf("error running HTTP controller")
			return err
		} else {
			time.Sleep(app.cfg.RevokeStunDuration)
		}
	}
}
