package controllers

import (
	"going/app/models"
)

type AbstractController struct {
	GorpController
}

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

func (c AbstractController) connected() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}
	if username, ok := c.Session["user"]; ok {
		user := c.getUser(username)
		c.RenderArgs["user"] = user
		return user
	}
	return nil
}
