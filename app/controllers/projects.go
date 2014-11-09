package controllers

import (
	"github.com/revel/revel"
	"going/app/models"
)

type Projects struct {
	AbstractController
}

/**
 * List all not expired projects in a page
 */
func (c Projects) List() revel.Result {
	c.connected()
	// get all not expired projects
	results, err := c.Txn.Select(models.Project{},
		`select * from Project WHERE expiration_date > NOW()`)
	if err != nil {
		panic(err)
	}

	var projects []*models.Project
	for _, r := range results {
		b := r.(*models.Project)
		b.Pledged = 0
		// get the pledged amount of the project
		results, err := c.Txn.SelectInt("select sum(amount) from Transaction WHERE project_id=?", b.Id)
		if err == nil {
			b.Pledged = results
		}
		projects = append(projects, b)
	}
	// display the page
	return c.Render(projects)

}
