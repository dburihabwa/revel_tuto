package controllers

import (
	"github.com/revel/revel"
)

type Projects struct {
	AbstractController
}

func (c Projects) List() revel.Result {
	c.connected()

	return c.Render()
}
