package services

import (
	"github.com/domgolonka/threatdefender/app"
	utils "github.com/domgolonka/threatdefender/pkg/utils/email"
	"time"
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
	_, domain := utils.Split(email)
	dom, err := utils.DomainAge(domain)
	if err != nil {
		app.Logger.Error(err)
	} else {
		// only display if domain age is accurate
		t1, err := time.Parse("1995-08-13T04:00:00Z", dom.CreatedDate)
		if err != nil {
			app.Logger.Error(err)
		}

		t2 := time.Now()
		days := t2.Sub(t1).Hours() / 24
		if days < 7 { // less than a week
			score += scoreCfg.Domain.Week
		} else if days < 30 {
			score += scoreCfg.Domain.Month
		} else if days < 365 {
			score += scoreCfg.Domain.Year
		} else if days >= 365 {
			score += scoreCfg.Domain.YearPlus
		}
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
	if app.PwnedKey != "" {
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
