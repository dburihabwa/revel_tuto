package controllers

import (
	"errors"
	"github.com/revel/revel"
	"going/app/models"
	"going/app/routes"
	"strconv"
	"time"
)

type Project struct {
	AbstractController
}

func (c Project) Index(id string) revel.Result {
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

	if c.Validation.HasErrors()  {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Project.AddProject())
	}

	project.OwnerId = user.Id
	project.CreationDate = time.Now()
	publicationDate, errPublicationParsing := MakeTime(publicationYear, publicationMonth, publicationDay)
	if errPublicationParsing != nil {
		panic(errPublicationParsing)
	}
	project.PublicationDate = publicationDate
	expirationDate, errExpirationParsing := MakeTime(expirationYear, expirationMonth, expirationDay)
	if errExpirationParsing != nil {
		panic(errExpirationParsing)
	}
	project.ExpirationDate = expirationDate
	if project.PublicationDate.After(project.ExpirationDate) {
		panic(errors.New("La date d'expiration est antérieure à la date de publication!"))
	}
	if project.CreationDate.After(project.ExpirationDate) {
		panic(errors.New("Le date d'expiration de la campagne est antérieure à la date d'aujourd'hui!"))
	}

    err := c.Txn.Insert(&project)
    if err != nil {
    	panic(err)
    }
    return c.Redirect(routes.Projects.List())
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
