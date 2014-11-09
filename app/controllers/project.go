package controllers

import (
	"github.com/revel/revel"
	"going/app/models"
	"going/app/routes"
	"time"
)

type Project struct {
	AbstractController
}

/**
 * Display the detail page of a project
 */
func (c Project) Index(Id int64) revel.Result {
	c.connected()

	// get the project
	obj, err := c.Txn.Get(models.Project{}, Id)
	if err != nil || obj == nil {
		// redirect the user if the project is not found
		return c.NotFound("Project not found.")
	}
	project := obj.(*models.Project)

	// count the number of contributors
	results, err := c.Txn.SelectInt("select count(DISTINCT user_id) from Transaction WHERE project_id=?", Id)
	var nbPledge int64 = 0
	if err == nil {
		nbPledge = results
	}

	// count the total amount pledged
	sumResults, err := c.Txn.SelectInt("select sum(amount) from Transaction WHERE project_id=?", Id)
	var pledged int64 = 0
	if err == nil {
		pledged = sumResults
	}

	// get the owner of the project
	ownerResult, err := c.Txn.Get(models.User{}, project.OwnerId)
	owner := ownerResult.(*models.User)

	// display the page
	return c.Render(project, nbPledge, pledged, owner)
}

/**
 * Reward a project
 */
func (c Project) Reward(transaction models.Transaction, amount int64, projectId int64) revel.Result {
	// checks if the current user is connected
	user := c.connected()
	if user == nil {
		return c.RenderText("Not conected")
	}
	// create the transaction
	transaction.UserId = user.Id
	transaction.ProjectId = projectId
	transaction.Amount = amount
	transaction.Date = time.Now()
	// insert the transaction
	err := c.Txn.Insert(&transaction)
	if err != nil {
		panic(err)
	}
	// redirect the user to the project page
	return c.Redirect(routes.Project.Index(projectId))
}

/**
 * Create the project page
 */
func (c Project) AddProject() revel.Result {
	// checks if the user is connected
	user := c.connected()
	if user == nil {
		// redirects the user to the login page
		return c.Redirect(routes.User.LoginPage())
	}
	return c.Render()
}

/**
 * Save the project in the database
 */
func (c Project) SaveProject(project models.Project, publicationDay string, publicationMonth string, publicationYear string,
	expirationDay string, expirationMonth string, expirationYear string) revel.Result {
	// checks if the user is connected
	user := c.connected()
	if user == nil {
		// redirects the user to the index page
		return c.Render(routes.Application.Index)
	}

	project.Validate(c.Validation)

	publicationDate, errPublicationParsing := MakeTime(publicationYear, publicationMonth, publicationDay)
	c.Validation.Required(errPublicationParsing == nil).Key("project.PublicationDate").
		Message("Date de publication invalide")

	expirationDate, errExpirationParsing := MakeTime(expirationYear, expirationMonth, expirationDay)
	c.Validation.Required(errExpirationParsing == nil).Key("project.ExpirationDate").
		Message("Date d'expiration invalide")

	c.Validation.Required(publicationDate.Before(expirationDate)).Key("project.PublicationDate").
		Message("La date d'expiration est antérieure à la date de publication!")

	now := time.Now()

	c.Validation.Required(now.Before(expirationDate)).Key("project.ExpirationDate").
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

	// insert the project in the database
	err := c.Txn.Insert(&project)
	if err != nil {
		panic(err)
	}
	return c.Redirect(routes.Offers.AddOffers(project.Id))
}

/**
 * Utility function used to create Time from text
 */
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

