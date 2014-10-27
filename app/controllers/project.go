package controllers

import (
	"github.com/revel/revel"
	"going/app/models"
)

type Project struct {
	AbstractController
}

func (c Project) Index(id string) revel.Result {
	return c.Render()
}

func (c Project) AddProject() revel.Result {
	return c.Render()
}

func (c Project) SaveProject(project models.Project) revel.Result {

	return nil
}
