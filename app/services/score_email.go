package services

import (
	"github.com/domgolonka/threatdefender/app"
)

func ScoreEmail(app *app.App, emailAddress string) (uint8, error) {
	var score uint8
	score = 0
	disposableEmail, err := app.DisableStore.FindByURL(emailAddress)
	if err != nil {
		app.Logger.Error(err)
		return score, err
	}
	spamEmail, err := app.SpamEmailStore.FindByURL(emailAddress)
	if err != nil {
		app.Logger.Error(err)
		return score, err
	}
	freeEmail, err := app.FreeEmailStore.FindByURL(emailAddress)
	if err != nil {
		app.Logger.Error(err)
		return score, err
	}

	isGeneric, err := GenericGetEmail(app, emailAddress)
	if err != nil {
		app.Logger.Error(err)
		return score, err
	}
	if disposableEmail != nil {
		score += 30
	}
	if spamEmail != nil {
		score += 50
	}
	if freeEmail != nil {
		score += 15
	}
	if *isGeneric {
		score += 10
	}
	if score > 100 {
		score = 100
	}

	return score, nil

}
