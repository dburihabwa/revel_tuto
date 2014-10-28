package models

import (
	"fmt"
	"github.com/revel/revel"
)

type Offer struct {
	Id          int64  `db:"id" json:"id"`
	ProjectId   int64  `db:"project_id" json:"project_id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Price       int64  `db:"price" json:"price"`
	Available   int    `db:"available" json:"available"`
}

func (u *Offer) String() string {
	return fmt.Sprintf("Offer(%s: %s)", u.Id, u.Title)
}

func (user *Offer) Validate(v *revel.Validation) {

	v.Check(user.Title,
		revel.Required{},
		revel.MaxSize{100},
	)

	v.Check(user.Description,
		revel.Required{},
		revel.MaxSize{1500},
	)
}
