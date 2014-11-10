package controllers

import (
	"github.com/revel/revel"
	"going/app/models"
	"going/app/routes"
	"strconv"
)

type Offers struct {
	AbstractController
}


/**
 * Create the project page
 */
func (c Offers) AddOffers(projectId int64) revel.Result {

	obj, err := c.Txn.Get(models.Project{}, projectId)
	if err != nil || obj == nil {
		// redirect the user if the project is not found
		return c.NotFound("Project not found.")
	}


	
	// checks if the user is connected
	user := c.connected()
	if user == nil {
		// redirects the user to the login page
		return c.Redirect(routes.User.LoginPage())
	}
	return c.Render(projectId)
}

/**
 * Save the project in the database
 */
func (c Offers) SaveOffers(offer models.Offer) revel.Result {
	// checks if the user is connected
	user := c.connected()
	if user == nil {
		// redirects the user to the index page
		return c.Render(routes.Application.Index)
	}

	offer.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Offers.AddOffers(offer.ProjectId))
	}

	// insert the project in the database
	err := c.Txn.Insert(&offer)
	if err != nil {
		panic(err)
	}
	return c.Redirect(routes.Offers.AddOffers(offer.ProjectId))
}


/**
 * Parse string to int
 */
func ParseStringToInt(str string) (int, error) {
	i64, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return 0, err
	}
	return int(i64), nil
}
