package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	figure "github.com/common-nighthawk/go-figure"
	"github.com/domgolonka/foretoken/app"
	"github.com/domgolonka/foretoken/app/data"
	"github.com/domgolonka/foretoken/config"
	"github.com/domgolonka/foretoken/lib/grpc/impl"
	"github.com/domgolonka/foretoken/server"
	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
)

func main() {
	var cmd string

	var appPath, _ = os.Getwd()
	configFilePath := appPath + "/config.yml"
	var cfg config.Config
	configFlag := flag.String("config", configFilePath, "the config file.")
	flag.Parse()

	// set the config file
	if configFlag != nil {
		configFilePath = *configFlag
	}
	err := configor.Load(&cfg, configFilePath)
	if err != nil {
		logrus.Info("\nsee: https://foretoken.domgolonka.com/docs/config/")
		logrus.Fatal(err)
	}
	// Default logger
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{}
	if cfg.Debug {
		logger.Level = logrus.DebugLevel
	}

	if len(os.Args) == 1 {
		cmd = "server"
	} else {
		cmd = os.Args[1]
	}
	switch {
	case cmd == "server":
		serve(&cfg, logger)
	case cmd == "migrate":
		migrate(&cfg, logger)
	default:
		_, err = os.Stderr.WriteString("unexpected invocation\n")
		if err != nil {
			return
		}
		usage(logger)
		os.Exit(2)
	}
}

func serve(cfg *config.Config, logger logrus.FieldLogger) {

	var (
		ch = make(chan bool)
	)
	myFigure := figure.NewFigure("Foretoken", "nancyj", true)
	myFigure.Print()

	newApp, err := app.NewApp(cfg, logger)
	if err != nil {
		logger.Fatal(err)
		return
	}

	impl.InitRPCService(newApp)
	go server.Server(newApp)
	//services.ServeRPC(newApp, ch)

	<-ch
}

func migrate(cfg *config.Config, logger logrus.FieldLogger) {
	logger.Info("Running migrations.")
	err := data.MigrateDB(cfg, nil)
	if err != nil {
		logger.Error(err)
	} else {
		logger.Info("Migrations complete.")
	}
}

func usage(logger logrus.FieldLogger) {
	exe := path.Base(os.Args[0])
	logger.Infof(fmt.Sprintf(`
Usage:
%s server  - run the server (default)
%s migrate - run migrations
`, exe, exe))
}
