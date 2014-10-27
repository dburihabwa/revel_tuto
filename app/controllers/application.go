package controllers

import (
	"github.com/revel/revel"
	"going/app/routes"
)

type Application struct {
	AbstractController
}

func (c Application) Index() revel.Result {
	if c.connected() != nil {
		return c.Redirect(routes.Projects.List())
	}
	c.Flash.Error("Please log in first")
	return c.Render()
}
