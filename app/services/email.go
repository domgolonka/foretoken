package services

import (
	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/app/entity"
	utils "github.com/domgolonka/threatdefender/pkg/utils/email"
)

type Email struct {
	email      string         `json:"-"`
	app        *app.App       `json:"-"`
	Disposable bool           `json:"disposable"`
	Free       bool           `json:"free"`
	Spam       bool           `json:"spam"`
	Generic    bool           `json:"generic"`
	CatchAll   bool           `json:"catchall"`
	Leaked     bool           `json:"leaked"`
	Valid      bool           `json:"valid"`
	Domain     *entity.Domain `json:"domain"`
}

func (e Email) EmailService() (*entity.EmailResponse, error) {
	emailsrv := &entity.EmailResponse{
		Disposable: false,
		Free:       false,
		RecentSpam: false,
		Valid:      true,
	}

	emailsrv.Leaked = e.Leaked
	emailsrv.Valid = e.Valid
	emailsrv.Disposable = e.Disposable
	emailsrv.RecentSpam = e.Spam
	emailsrv.CatchAll = e.CatchAll
	emailsrv.Free = e.Free
	emailsrv.Domain = e.Domain
	emailsrv.Generic = e.Generic

	score, err := e.ScoreEmail()
	if err != nil {
		e.app.Logger.Error(err)
	} else {
		emailsrv.Score = score
	}

	emailsrv.Success = true
	return emailsrv, nil
}

func (e Email) Calculate(app *app.App, email string) {
	e.app = app
	e.email = email
	disposable, err := app.DisableStore.FindByDomain(email)
	if err != nil {
		app.Logger.Error(err)
	}
	if disposable != nil {
		e.Disposable = true
	}
	freeEmail, err := app.FreeEmailStore.FindByDomain(email)
	if err != nil {
		app.Logger.Error(err)
	}
	if freeEmail != nil {
		e.Free = true
	}
	spamEmail, err := app.SpamEmailStore.FindByDomain(email)
	if err != nil {
		app.Logger.Error(err)
	}
	if spamEmail != nil {
		e.Spam = true
	}
	genericEmail, err := GenericGetEmail(app, email)
	if err != nil {
		app.Logger.Error(err)
	}
	e.Generic = *genericEmail
	err = utils.ValidateEmail(app, email)
	if err != nil {
		e.Valid = false
	}
	// only use catch all if smtp is enabled
	used, err := utils.CatchAll(app, email)
	if err != nil && used {
		e.CatchAll = false
	}
	_, domain := utils.Split(email)
	dom, err := utils.DomainAge(domain)
	if err != nil {
		app.Logger.Error(err)
	} else {
		e.Domain = dom
	}
	if app.PwnedKey != "" {
		leaked, err := utils.Leaked(app, email, "")
		if err != nil {
			app.Logger.Error(err)
		}
		e.Leaked = *leaked
	}
}
