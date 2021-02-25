package app

import (
	"github.com/domgolonka/threatdefender/pkg/utils/ip"
	"time"

	"github.com/domgolonka/threatdefender/app/data"
	"github.com/domgolonka/threatdefender/config"
	"github.com/domgolonka/threatdefender/lib/scrapers/email/disposable"
	"github.com/domgolonka/threatdefender/lib/scrapers/email/free"
	spamemail "github.com/domgolonka/threatdefender/lib/scrapers/email/spam"
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
	FreeEmailStore      data.FreeEmailStore
	ProxyGenerator      *proxy.ProxyGenerator
	VPNGenerator        *vpn.VPN
	DisposableGenerator *disposable.Disposable
	SpamGenerator       *spam.Spam
	SpamEmailGenerator  *spamemail.SpamEmail
	TorGenerator        *tor.Tor
	FreeEmailGenerator  *free.Free
	Maxmind             *ip.Maxmind
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
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "error cannot ping to database")
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
	freeEmailStore, err := data.NewFreeEmailStore(db)
	if err != nil {
		return nil, errors.Wrap(err, "NewFreeEmailStore")
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

	var maxmind *ip.Maxmind
	if cfg.APIKeys.Maxmind != "" {
		maxmind = ip.NewMaxmind(cfg.APIKeys.Maxmind)
	}

	proxygen := proxy.New(proxyStore, cfg.Proxy.Workers, time.Duration(cfg.Proxy.CacheDurationMinutes), logger)
	vpngen := vpn.NewVPN(vpnStore, logger)
	torgen := tor.NewTor(torStore, logger)
	spamgen := spam.NewSpam(spamStore, logger)
	freeEmailGen := free.NewFreeEmail(freeEmailStore, logger)
	disgen := disposable.NewDisposable(disposableStore, logger)
	spamemailgen := spamemail.NewSpamEmail(spamEmailStore, logger)

	return &App{
		// Provide access to root DB - useful when extending AccountStore functionality
		Config:   cfg,
		Reporter: errorReporter,
		Logger:   logger,
		// store
		ProxyStore:     proxyStore,
		VpnStore:       vpnStore,
		DisableStore:   disposableStore,
		FreeEmailStore: freeEmailStore,
		SpamStore:      spamStore,
		SpamEmailStore: spamEmailStore,
		TorStore:       torStore,
		// generator
		ProxyGenerator:      proxygen,
		VPNGenerator:        vpngen,
		TorGenerator:        torgen,
		SpamGenerator:       spamgen,
		DisposableGenerator: disgen,
		SpamEmailGenerator:  spamemailgen,
		FreeEmailGenerator:  freeEmailGen,
		Maxmind:             maxmind,
	}, nil
}
