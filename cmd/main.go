package main

import (
	"fmt"
	"github.com/domgolonka/threatscraper/app"
	"github.com/domgolonka/threatscraper/app/config"
	"github.com/domgolonka/threatscraper/app/data"
	"github.com/domgolonka/threatscraper/server"
	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

// VERSION is a value injected at build time with ldflags
var VERSION string

func main() {
	var cmd string
	var appPath, _ = os.Getwd()
	configFilePath := appPath + "/app/config/config.dev.yml"
	var cfg config.Config
	err := configor.Load(&cfg, configFilePath)
	if err != nil {
		fmt.Println(err)
		fmt.Println("\nsee: https://github.com/github.com/domgolonka/threatscraper/blob/master/docs/config.md")
		return
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
		os.Stderr.WriteString(fmt.Sprintf("unexpected invocation\n"))
		usage()
		os.Exit(2)
	}
}

func serve(cfg config.Config) {
	fmt.Println(fmt.Sprintf("~*~ Defend.Export v%s ~*~", VERSION))

	// Default logger
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	if cfg.Debug {
		logger.Level = logrus.DebugLevel
	}
	logger.Out = os.Stdout

	newApp, err := app.NewApp(cfg, logger)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fmt.Sprintf("THREAT_SCRAPER_URL: %s", cfg.Rooturl))
	fmt.Println(fmt.Sprintf("PORT: %d", cfg.ServerPort))
	if newApp.Config.PublicPort != 0 {
		fmt.Println(fmt.Sprintf("PUBLIC_PORT: %d", newApp.Config.PublicPort))
	}

	server.Server(newApp)
}

func migrate(cfg config.Config) {
	fmt.Println("Running migrations.")
	err := data.MigrateDB(cfg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Migrations complete.")
	}
}

func usage() {
	exe := path.Base(os.Args[0])
	fmt.Println(fmt.Sprintf(`
Usage:
%s server  - run the server (default)
%s migrate - run migrations
`, exe, exe))
}
