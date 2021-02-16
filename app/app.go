package app

import (
	"time"

	spamemail "github.com/domgolonka/threatdefender/lib/scrapers/email/spam"

	"github.com/domgolonka/threatdefender/app/data"
	"github.com/domgolonka/threatdefender/config"
	"github.com/domgolonka/threatdefender/lib/scrapers/email/disposable"
	"github.com/domgolonka/threatdefender/lib/scrapers/ip/proxy"
	"github.com/domgolonka/threatdefender/lib/scrapers/ip/spam"
	"github.com/domgolonka/threatdefender/lib/scrapers/ip/tor"
	"github.com/domgolonka/threatdefender/lib/scrapers/ip/vpn"
	"github.com/domgolonka/threatdefender/ops"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type App struct {
	ProxyList           []string
	SpamList            []string
	VPNList             []string
	DCList              []string
	EmailDisposalList   []string
	EmailFreeList       []string
	Logger              logrus.FieldLogger
	Reporter            ops.ErrorReporter
	Config              config.Config
	ProxyStore          data.ProxyStore
	VpnStore            data.VpnStore
	DisableStore        data.DisposableStore
	SpamStore           data.SpamStore
	SpamEmailStore      data.SpamEmailStore
	TorStore            data.TorStore
	ProxyGenerator      *proxy.ProxyGenerator
	VPNGenerator        *vpn.VPN
	DisposableGenerator *disposable.Disposable
	SpamGenerator       *spam.Spam
	SpamEmailGenerator  *spamemail.SpamEmail
	TorGenerator        *tor.Tor
}

func NewApp(cfg config.Config, logger logrus.FieldLogger) (*App, error) {
	reporter := ops.Log
	if cfg.ErrorReporter.Default == "airbreak" {
		reporter = ops.Airbrake
	} else if cfg.ErrorReporter.Default == "sentry" {
		reporter = ops.Sentry
	}
	errorReporter, err := ops.NewErrorReporter(cfg.ErrorReporter.AirbrakeCredentials, reporter, logger)
	if err != nil {
		return nil, errors.Wrap(err, "ReporterError")
	}

	db, err := data.NewDB(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "data.NewDB")
	}

	proxyStore, err := data.NewProxyStore(db)
	if err != nil {
		return nil, errors.Wrap(err, "NewProxyStore")
	}
	vpnStore, err := data.NewVpnStore(db)
	if err != nil {
		return nil, errors.Wrap(err, "NewVPNStore")
	}
	disposableStore, err := data.NewDisposableStore(db)
	if err != nil {
		return nil, errors.Wrap(err, "NewDisposableStore")
	}
	spamStore, err := data.NewSpamStore(db)
	if err != nil {
		return nil, errors.Wrap(err, "NewSpamStore")
	}
	spamEmailStore, err := data.NewSpamEmailStore(db)
	if err != nil {
		return nil, errors.Wrap(err, "NewSpamEmailStore")
	}
	torStore, err := data.NewTorStore(db)
	if err != nil {
		return nil, errors.Wrap(err, "NewSpamStore")
	}

	proxygen := proxy.New(proxyStore, cfg.Proxy.Workers, time.Duration(cfg.Proxy.CacheDurationMinutes), logger)
	vpngen := vpn.NewVPN(vpnStore, logger)
	disgen := disposable.NewDisposable(disposableStore, logger)
	spamgen := spam.NewSpam(spamStore, logger)
	spamemailgen := spamemail.NewSpamEmail(spamEmailStore, logger)
	torgen := tor.NewTor(torStore, logger)

	return &App{
		// Provide access to root DB - useful when extending AccountStore functionality
		Config:              cfg,
		Reporter:            errorReporter,
		Logger:              logger,
		ProxyStore:          proxyStore,
		VpnStore:            vpnStore,
		DisableStore:        disposableStore,
		SpamStore:           spamStore,
		SpamEmailStore:      spamEmailStore,
		TorStore:            torStore,
		ProxyGenerator:      proxygen,
		VPNGenerator:        vpngen,
		DisposableGenerator: disgen,
		SpamGenerator:       spamgen,
		SpamEmailGenerator:  spamemailgen,
		TorGenerator:        torgen,
	}, nil
}
