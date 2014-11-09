package controllers

import (
	"going/app/models"
)

type AbstractController struct {
	GorpController
}

/**
 * Get a user by his username
 */
func (c AbstractController) getUser(username string) *models.User {
	users, err := c.Txn.Select(models.User{}, `select * from User where Username = ?`, username)
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		return nil
	}
	return users[0].(*models.User)
}

/**
 * Checks if the current client is connected
 */
func (c AbstractController) connected() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}
	// get the user form the session
	if username, ok := c.Session["user"]; ok {
		// get the user
		user := c.getUser(username)
		// add the current user to the view parameters
		c.RenderArgs["user"] = user
		return user
	}
	return nil
}
