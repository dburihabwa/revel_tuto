package models

import (
	"fmt"
	"github.com/revel/revel"
	"time"
)

type Project struct {
	Id              int64     `db:"id" json:"id"`
	Title           string    `db:"title" json:"title"`
	Description     string    `db:"description" json:"description"`
	OwnerId         int64     `db:"owner" json:"owner"`
	Amount          int64     `db:"amount" json: "amount"`
	Pledged         int64     `db:"-"`
	PublicationDate time.Time `db:"pubication_date" json:"pubication_date"`
	CreationDate    time.Time `db:"creation_date" json:"creation_date"`
	ExpirationDate  time.Time `db:"expiration_date" json:"expiration_date"`
}

func (p *Project) String() string {
	return fmt.Sprintf("Project(%s: %s)", p.Id, p.Title)
}

func (project *Project) Validate(v *revel.Validation) {

	v.Check(project.Title,
		revel.Required{},
		revel.MaxSize{100},
	)

	v.Check(project.Description,
		revel.Required{},
		revel.MaxSize{1500},
	)

	v.Check(int(project.Amount),
		revel.Required{},
		revel.Range{revel.Min{0}, revel.Max{1000000000}},
	).Message("Le montant doit être compris entre 0 et 1 000 000 000.")
}
