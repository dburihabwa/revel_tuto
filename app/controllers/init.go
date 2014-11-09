package controllers

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/revel/revel"
	"github.com/russross/blackfriday"
	T "html/template"
	"time"
)

func init() {
	revel.OnAppStart(InitDB)
	revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	revel.InterceptMethod(User.AddUser, revel.BEFORE)
	revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)

	/**
	 * Template functions
	 */

	// count the number of days between today and date
	revel.TemplateFuncs["dayLeft"] = func(date time.Time) int {
		today := time.Now()
		return int(date.Sub(today).Hours() / 24)
	}

	// perform the difference between a and b
	revel.TemplateFuncs["diff"] = func(a int64, b int64) int64 {
		return a - b
	}

	// perform the percent
	revel.TemplateFuncs["percent"] = func(total int64, current int64) int64 {
		val := int64((float64(current) / float64(total)) * 100)
		if val < 0 {
			return 0
		}
		return val
	}
	// perform the percent with and limit the percent to a maximum value
	revel.TemplateFuncs["percentMax"] = func(total int64, current int64, max int64) int64 {
		val := int64((float64(current) / float64(total)) * 100)
		if val < 0 {
			return 0
		}
		if val > max {
			return max
		}
		return val
	}
	// convert markdown to HTML
	revel.TemplateFuncs["md"] = func(markdown string) T.HTML {
		unsafe := blackfriday.MarkdownCommon([]byte(markdown))
		html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
		return T.HTML(html)
	}
}
