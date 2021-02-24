package services

import (
	"github.com/domgolonka/threatdefender/app"
	utils "github.com/domgolonka/threatdefender/pkg/utils/email"
)

func ScoreEmail(app *app.App, email string) (int8, error) {
	var score int8
	scoreCfg := app.Config.Email.Score
	score = 0
	disposableEmail, err := app.DisableStore.FindByDomain(email)
	if err != nil {
		app.Logger.Error(err)
		return score, err
	}
	spamEmail, err := app.SpamEmailStore.FindByDomain(email)
	if err != nil {
		app.Logger.Error(err)
		return score, err
	}
	freeEmail, err := app.FreeEmailStore.FindByDomain(email)
	if err != nil {
		app.Logger.Error(err)
		return score, err
	}

	err = utils.ValidateEmail(app, email)
	// is not a valid email
	if err != nil {
		score += scoreCfg.Valid.Yes
	} else {
		score += scoreCfg.Valid.No
	}
	// only use catch all if smtp is enabled
	used, err := utils.CatchAll(app, email)
	if err != nil && used {
		score += scoreCfg.CatchAll.Yes
	} else {
		score += scoreCfg.CatchAll.No
	}

	// only use catch all if smtp is enabled
	leaked, err := utils.Leaked(app, email, "")
	if err != nil {
		app.Logger.Error(err)
		return score, err
	}
	if *leaked {
		score += scoreCfg.Leaked.Yes
	} else {
		score += scoreCfg.Leaked.No
	}

	isGeneric, err := GenericGetEmail(app, email)
	if err != nil {
		app.Logger.Error(err)
		return score, err
	}
	if disposableEmail != nil {
		score += scoreCfg.Disposable.Yes
	} else {
		score += scoreCfg.Disposable.No
	}
	if spamEmail != nil {
		score += scoreCfg.Spam.Yes
	} else {
		score += scoreCfg.Spam.No
	}
	if freeEmail != nil {
		score += scoreCfg.Free.Yes
	} else {
		score += scoreCfg.Free.No
	}
	if *isGeneric {
		score += scoreCfg.Generic.Yes
	} else {
		score += scoreCfg.Generic.No
	}
	if score > 100 {
		score = 100
	} else if score < 0 {
		score = 0
	}

	return score, nil

}
