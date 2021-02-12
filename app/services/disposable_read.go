package services

import (
	"io/ioutil"

	"github.com/domgolonka/threatscraper/config"
	"github.com/pkg/errors"
)

func DisposableRead(cfg *config.Config) ([]byte, error) {
	b, err := ioutil.ReadFile(cfg.DisposableFile)
	return b, errors.Wrap(err, "READFILE")

}
