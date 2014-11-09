package controllers

import (
	"github.com/revel/revel"
	"going/app/models"
)

type Application struct {
	AbstractController
}

/**
 * Render the Index page
 */
func (c Application) Index() revel.Result {
	c.connected()

	// Get popular projects
	results, err := c.Txn.Select(models.Project{},
		`select DISTINCT p.* from Project p INNER JOIN transaction t on t.project_id=p.id WHERE p.expiration_date > NOW() GROUP BY p.id ORDER BY count(t.user_id) DESC LIMIT 3`)
	if err != nil {
		panic(err)
	}

	var projects []*models.Project
	for _, r := range results {
		b := r.(*models.Project)
		b.Pledged = 0
		// get the pledged amount
		results, err := c.Txn.SelectInt("select sum(amount) from transaction WHERE project_id=?", b.Id)
		if err == nil {
			b.Pledged = results
		}
		projects = append(projects, b)
	}

	return c.Render(projects)
}
