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

func (c Project) Index(id string) revel.Result {
	return c.Render()
}

func (c Project) AddProject() revel.Result {
	return c.Render()
}

func (c Project) SaveProject(project models.Project) revel.Result {
	project.CreationDate = time.Now()
    err := c.Txn.Insert(&project)
    if err != nil {
    	panic(err)
    }
    return c.Redirect(routes.Projects.List())
}
