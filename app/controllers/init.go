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

	revel.TemplateFuncs["diff"] = func(a int64, b int64) int64 {
		return a - b
	}

	revel.TemplateFuncs["procent"] = func(a int64, b int64) float64 {
		return (float64(b) / float64(a)) * 100
	}
}
