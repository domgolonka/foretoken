package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	services "github.com/domgolonka/threatdefender/lib/grpc"
	"github.com/domgolonka/threatdefender/lib/grpc/impl"

	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/app/data"
	"github.com/domgolonka/threatdefender/config"
	"github.com/domgolonka/threatdefender/server"
	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
)

// VERSION is a value injected at build time with ldflags
var VERSION string

func main() {
	var cmd string
	var appPath, _ = os.Getwd()
	configFilePath := appPath + "/config/config.dev.yml"
	var cfg config.Config
	configFlag := flag.String("config", configFilePath, "the config file.")
	flag.Parse()
	// set the config file
	if configFlag != nil {
		configFilePath = *configFlag
	}
	err := configor.Load(&cfg, configFilePath)
	if err != nil {
		logrus.Info("\nsee: https://github.com/domgolonka/threatdefender/blob/master/docs/config.md")
		logrus.Fatal(err)
	}

	if len(os.Args) == 1 {
		cmd = "server"
	} else {
		cmd = os.Args[1]
	}

	if cmd == "server" {
		serve(cfg)
	} else if cmd == "migrate" {
		migrate(cfg)
	} else {
		os.Stderr.WriteString("unexpected invocation\n")
		usage()
		os.Exit(2)
	}
}

func serve(cfg config.Config) {

	var (
		ch = make(chan bool)
	)

	// Default logger
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	if cfg.Debug {
		logger.Level = logrus.DebugLevel
	}
	logger.Out = os.Stdout

	logger.Infof(fmt.Sprintf("~*~ ThreatDefender v%s ~*~", VERSION))

	newApp, err := app.NewApp(cfg, logger)
	if err != nil {
		logger.Fatal(err)
		return
	}

	impl.InitRPCService(newApp)
	go server.Server(newApp)
	services.ServeRPC(newApp, ch, cfg.GRPCPort)

	<-ch
}

func migrate(cfg config.Config) {
	logrus.Info("Running migrations.")
	err := data.MigrateDB(cfg)
	if err != nil {
		logrus.Error(err)
	} else {
		logrus.Info("Migrations complete.")
	}
}

func usage() {
	exe := path.Base(os.Args[0])
	logrus.Info(fmt.Sprintf(`
Usage:
%s server  - run the server (default)
%s migrate - run migrations
`, exe, exe))
}
