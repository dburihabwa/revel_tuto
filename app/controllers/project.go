package controllers

import (
	"github.com/revel/revel"
	"going/app/models"
	"going/app/routes"
	"strconv"
	"time"
)

type Project struct {
	AbstractController
}

func (c Project) Index(Id int64) revel.Result {
	return c.Render()
}

func (c Project) AddProject() revel.Result {
	user := c.connected()
	if user == nil {
		return c.Render(routes.Application.Index)
	}
	return c.Render()
}

func (c Project) SaveProject(project models.Project, publicationDay string, publicationMonth string, publicationYear string,
	expirationDay string, expirationMonth string, expirationYear string) revel.Result {
	user := c.connected()
	if user == nil {
		return c.Render(routes.Application.Index)
	}

	project.Validate(c.Validation)

	publicationDate, errPublicationParsing := MakeTime(publicationYear, publicationMonth, publicationDay)
	c.Validation.Required(errPublicationParsing == nil).Key("project.PublicationDate").
		Message("Date de publication invalide")

	expirationDate, errExpirationParsing := MakeTime(expirationYear, expirationMonth, expirationDay)
	c.Validation.Required(errExpirationParsing == nil).Key("project.ExpirationDate").
		Message("Date d'expiration invalide")

	c.Validation.Required(publicationDate.Before(project.ExpirationDate)).Key("project.PublicationDate").
		Message("La date d'expiration est antérieure à la date de publication!")

	now := time.Now()

	c.Validation.Required(now.Before(project.ExpirationDate)).Key("project.ExpirationDate").
		Message("Le date d'expiration de la campagne est antérieure à la date d'aujourd'hui!")

	project.OwnerId = user.Id
	project.CreationDate = now
	project.PublicationDate = publicationDate
	project.ExpirationDate = expirationDate

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Project.AddProject())
	}

	err := c.Txn.Insert(&project)
	if err != nil {
		panic(err)
	}
	return c.Redirect(routes.Project.Index(project.Id))
}

func MakeTime(yearString string, monthString string, dayString string) (time.Time, error) {
	day, errParseDay := ParseStringToInt(dayString)
	if errParseDay != nil {
		return time.Now(), errParseDay
	}
	month, errParseMonth := ParseStringToInt(monthString)
	if errParseMonth != nil {
		return time.Now(), errParseMonth
	}
	year, errParseYear := ParseStringToInt(yearString)
	if errParseYear != nil {
		return time.Now(), errParseYear
	}
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.FixedZone("Europe/Paris", 0)), nil
}

func ParseStringToInt(str string) (int, error) {
	i64, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return 0, err
	}
	return int(i64), nil
}
