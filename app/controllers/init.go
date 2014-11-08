package controllers

import (
	"github.com/revel/revel"
	"time"
)

func init() {
	revel.OnAppStart(InitDB)
	revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	revel.InterceptMethod(User.AddUser, revel.BEFORE)
	revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)

	revel.TemplateFuncs["dayLeft"] = func(date time.Time) int {
		today := time.Now()
		return int(date.Sub(today).Hours() / 24)
	}
}
