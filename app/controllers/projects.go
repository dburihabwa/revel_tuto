package controllers

import (
	"github.com/revel/revel"
	"going/app/models"
)

type Projects struct {
	AbstractController
}

func (c Projects) List() revel.Result {
	c.connected()

	results, err := c.Txn.Select(models.Project{},
		`select * from Project`)
	if err != nil {
		panic(err)
	}

	var projects []*models.Project
	for _, r := range results {
		b := r.(*models.Project)
		b.Pledged = 0
		projects = append(projects, b)
	}

	return c.Render(projects)

}
