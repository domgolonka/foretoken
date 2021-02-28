package services

import (
	"strings"

	"github.com/domgolonka/foretoken/app"
	"github.com/domgolonka/foretoken/app/entity"
	utils "github.com/domgolonka/foretoken/pkg/utils/email"
)

type Email struct {
	email      string
	app        *app.App
	Disposable bool           `json:"disposable"`
	Free       bool           `json:"free"`
	Spam       bool           `json:"spam"`
	Generic    bool           `json:"generic"`
	CatchAll   bool           `json:"catchall"`
	Leaked     bool           `json:"leaked"`
	Valid      bool           `json:"valid"`
	Domain     *entity.Domain `json:"domain"`
}

func (e *Email) EmailService() (*entity.EmailResponse, error) {
	emailSrv := &entity.EmailResponse{
		Disposable: false,
		Free:       false,
		RecentSpam: false,
		Valid:      true,
	}

	emailSrv.Leaked = e.Leaked
	emailSrv.Valid = e.Valid
	emailSrv.Disposable = e.Disposable
	emailSrv.RecentSpam = e.Spam
	emailSrv.CatchAll = e.CatchAll
	emailSrv.Free = e.Free
	emailSrv.Domain = e.Domain
	emailSrv.Generic = e.Generic

	score, err := e.ScoreEmail()
	if err != nil {
		e.app.Logger.Error(err)
	} else {
		emailSrv.Score = score
	}

	emailSrv.Success = true
	return emailSrv, nil
}

func (e *Email) Calculate(app *app.App, email string) {
	e.app = app
	e.email = strings.ToLower(email)
	disposable, err := e.app.DisableStore.FindByDomain(email)
	if err != nil {
		e.app.Logger.Error(err)
	}
	if disposable != nil {
		e.Disposable = true
	}
	freeEmail, err := e.app.FreeEmailStore.FindByDomain(email)
	if err != nil {
		e.app.Logger.Error(err)
	}
	if freeEmail != nil {
		e.Free = true
	}
	spamEmail, err := e.app.SpamEmailStore.FindByDomain(email)
	if err != nil {
		e.app.Logger.Error(err)
	}
	if spamEmail != nil {
		e.Spam = true
	}
	genericEmail, err := GenericGetEmail(e.app, email)
	if err != nil {
		e.app.Logger.Error(err)
	}
	e.Generic = *genericEmail
	err = utils.ValidateEmail(e.app, email)
	if err != nil {
		e.Valid = false
	}
	// only use catch all if smtp is enabled
	used, err := utils.CatchAll(e.app, email)
	if err != nil && used {
		e.CatchAll = false
	}
	_, domain := utils.Split(email)
	dom, err := utils.DomainAge(domain)
	if err != nil {
		e.app.Logger.Error(err)
	} else {
		e.Domain = dom
	}
	if e.app.PwnedKey != "" {
		leaked, err := utils.Leaked(e.app, email, "")
		if err != nil {
			e.app.Logger.Error(err)
		}
		e.Leaked = *leaked
	}
}
