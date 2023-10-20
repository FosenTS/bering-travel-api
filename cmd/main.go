package main

import (
	"bering-travel-api/internal/application"
	"bering-travel-api/internal/application/config"
	"context"
	"flag"
	"fmt"
	mylogging "github.com/FosenTS/pkg/mylogger"
	sq "github.com/Masterminds/squirrel"
	"github.com/lann/builder"
	"log"
	"time"
)

var modeFlag *string

func main() {
	time.Local = time.UTC
	modeFlag = flag.String("mode", config.Mode.Local(), "specifies whether program is for dev or prod or local usage, default is local")

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		if cancel != nil {
			cancel()
		}
	}()
	flag.Parse()
	mode := *modeFlag
	switch mode {
	case config.Mode.Local():
		log.Println("Local mode!")
	case config.Mode.Dev():
		log.Println("Dev mode!")
	case config.Mode.Prod():
		log.Println("Prod mode!")
	default:
		log.Fatalln("invalid mode")
	}
	config.LoadEnv(mode)

	mylogging.ConfigureLogrus()
	rootLogger := mylogging.GetLogger()

	rootLogger.Infoln("Let's go!")
	rootLogger.Infoln("Creating app...")

	sq.StatementBuilder = sq.StatementBuilderType(builder.EmptyBuilder).PlaceholderFormat(sq.Dollar)

	app, err := application.NewApp(ctx, rootLogger.WithContext(ctx), nil)
	if err != nil {
		cancel()
		time.Sleep(time.Microsecond * 300)
		rootLogger.Fatalf(fmt.Sprintf("fatal creating app: %s\n", err))
	}
	rootLogger.Infoln("Starting app...")
	err = app.Run(ctx)
	if err != nil {
		cancel()
		time.Sleep(time.Microsecond * 300)
		log.Fatalf("fatal starting or while running app: %s\n", err)
	}
}
