package services

import (
	"github.com/domgolonka/threatdefender/app"
	utils "github.com/domgolonka/threatdefender/pkg/utils/email"
)

func ScoreEmail(app *app.App, email string) (uint8, error) {
	var score uint8
	score = 0
	disposableEmail, err := app.DisableStore.FindByEmail(email)
	if err != nil {
		app.Logger.Error(err)
		return score, err
	}
	spamEmail, err := app.SpamEmailStore.FindByEmail(email)
	if err != nil {
		app.Logger.Error(err)
		return score, err
	}
	freeEmail, err := app.FreeEmailStore.FindByEmail(email)
	if err != nil {
		app.Logger.Error(err)
		return score, err
	}

	err = utils.ValidateEmail(app, email)
	// is not a valid email
	if err != nil {
		score += 8
	}
	// only use catch all if smtp is enabled
	used, err := utils.CatchAll(app, email)
	if err != nil && used {
		score += 30
	}

	isGeneric, err := GenericGetEmail(app, email)
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
