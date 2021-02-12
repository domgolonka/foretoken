package services

import (
	"github.com/domgolonka/threatscraper/app/config"
	"github.com/pkg/errors"
	"io/ioutil"
)

func DisposableRead(cfg *config.Config) ([]byte, error) {
	b, err := ioutil.ReadFile(cfg.DisposableFile)
	return b, errors.Wrap(err, "READFILE")

}
