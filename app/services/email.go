package services

import (
	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/app/entity"
	utils "github.com/domgolonka/threatdefender/pkg/utils/email"
)

func EmailService(app *app.App, email string) (*entity.EmailResponse, error) {
	emailsrv := &entity.EmailResponse{
		Disposable: false,
		Free:       false,
		RecentSpam: false,
		Valid:      true,
	}
	disposable, err := app.DisableStore.FindByDomain(email)
	if err != nil {
		app.Logger.Error(err)
	}
	if disposable != nil {
		emailsrv.Disposable = true
	}
	freeEmail, err := app.FreeEmailStore.FindByDomain(email)
	if err != nil {
		app.Logger.Error(err)
	}
	if freeEmail != nil {
		emailsrv.Free = true
	}
	spamEmail, err := app.SpamEmailStore.FindByDomain(email)
	if err != nil {
		app.Logger.Error(err)
	}
	if spamEmail != nil {
		emailsrv.RecentSpam = true
	}
	genericEmail, err := GenericGetEmail(app, email)
	if err != nil {
		app.Logger.Error(err)
	}
	emailsrv.Generic = *genericEmail
	err = utils.ValidateEmail(app, email)
	if err != nil {
		emailsrv.Valid = false
	}
	// only use catch all if smtp is enabled
	used, err := utils.CatchAll(app, email)
	if err != nil && used {
		emailsrv.CatchAll = false
	}
	score, err := ScoreEmail(app, email)
	if err != nil {
		app.Logger.Error(err)
	} else {
		emailsrv.Score = score
	}
	_, domain := utils.Split(email)
	dom, err := utils.DomainAge(domain)
	if err != nil {
		app.Logger.Error(err)
	} else {
		emailsrv.Domain = dom
	}

	if app.PwnedKey != "" {
		leaked, err := utils.Leaked(app, email, "")
		if err != nil {
			app.Logger.Error(err)
		}
		emailsrv.Leaked = *leaked
	}
	return emailsrv, nil
}
