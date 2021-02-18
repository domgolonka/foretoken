package services

import (
	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/app/entity"
)

func EmailService(app *app.App, email string) (*entity.EmailResponse, error) {
	emailsrv := &entity.EmailResponse{
		Disposable: false,
		Free:       false,
		RecentSpam: false,
		Valid:      true,
	}
	disposable, err := app.DisableStore.FindByEmail(email)
	if err != nil {
		app.Logger.Error(err)
	}
	if disposable != nil {
		emailsrv.Disposable = true
	}
	freeEmail, err := app.FreeEmailStore.FindByEmail(email)
	if err != nil {
		app.Logger.Error(err)
	}
	if freeEmail != nil {
		emailsrv.Free = true
	}
	spamEmail, err := app.SpamEmailStore.FindByEmail(email)
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
	err = ValidateEmail(app, email)
	if err != nil {
		emailsrv.RecentSpam = false
	}
	score, err := ScoreEmail(app, email)
	if err != nil {
		app.Logger.Error(err)
	} else {
		emailsrv.Score = score
	}
	_, domain := split(email)
	age, err := domainAge(domain)
	if err != nil {
		app.Logger.Error(err)
	} else {
		emailsrv.DomainAge = age
	}
	return emailsrv, nil
}
