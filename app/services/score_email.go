package services

import (
	"time"
)

func (e *Email) ScoreEmail() (int8, error) {
	var score int8
	scoreCfg := e.app.Config.Email.Score
	score = 0

	if e.Domain != nil {
		// only display if domain age is accurate
		t1, err := time.Parse("1996-03-27T05:00:00Z", e.Domain.CreatedDate)
		if err != nil {
			e.app.Logger.Error(err)
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

	// is not a valid email
	if e.Valid {
		score += scoreCfg.Valid.Yes
	} else {
		score += scoreCfg.Valid.No
	}

	if e.CatchAll {
		score += scoreCfg.CatchAll.Yes
	} else {
		score += scoreCfg.CatchAll.No
	}

	// only use Pwned if key is set
	if e.app.PwnedKey != "" {
		if e.Leaked {
			score += scoreCfg.Leaked.Yes
		} else {
			score += scoreCfg.Leaked.No
		}
	}

	if e.Disposable {
		score += scoreCfg.Disposable.Yes
	} else {
		score += scoreCfg.Disposable.No
	}
	if e.Spam {
		score += scoreCfg.Spam.Yes
	} else {
		score += scoreCfg.Spam.No
	}
	if e.Free {
		score += scoreCfg.Free.Yes
	} else {
		score += scoreCfg.Free.No
	}
	if e.Generic {
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
