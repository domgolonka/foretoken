package app

import (
	"strconv"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func jobs(logger logrus.FieldLogger, app *App) error {

	c := cron.New(cron.WithChain(
		cron.Recover(cron.VerbosePrintfLogger(logger)), // or use cron.DefaultLogger
	))

	_, err := c.AddFunc("@every "+strconv.Itoa(app.Config.Crontab.Tor)+"h", func() { app.TorGenerator.Run(app.Config.Crontab.Tor) })
	if err != nil {
		return err
	}
	_, err = c.AddFunc("@every "+strconv.Itoa(app.Config.Crontab.Spam)+"h", func() { app.SpamGenerator.Run(app.Config.Crontab.Spam) })
	if err != nil {
		return err
	}
	_, err = c.AddFunc("@every "+strconv.Itoa(app.Config.Crontab.VPN)+"h", func() { app.VPNGenerator.Run(app.Config.Crontab.VPN) })
	if err != nil {
		return err
	}
	_, err = c.AddFunc("@every "+strconv.Itoa(app.Config.Crontab.Proxy)+"h", func() { app.ProxyGenerator.Run(app.Config.Proxy.Workers, app.Config.Crontab.Proxy) })
	if err != nil {
		return err
	}
	_, err = c.AddFunc("@every "+strconv.Itoa(app.Config.Crontab.Proxy)+"h", func() { app.ProxyGenerator.Run(app.Config.Proxy.Workers, app.Config.Crontab.Proxy) })
	if err != nil {
		return err
	}
	_, err = c.AddFunc("@every "+strconv.Itoa(app.Config.Crontab.Maxmind)+"h", func() {
		err = app.Maxmind.DownloadAndUpdate()
		if err != nil {
			app.Logger.Error(err)
		}
	})
	if err != nil {
		return err
	}
	c.Start()
	return nil
}
