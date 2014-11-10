package models

import (
	"fmt"
	"github.com/revel/revel"
)

type Offer struct {
	Id          int64  `db:"id" json:"id"`
	ProjectId   int64  `db:"project_id" json:"project_id"`
	Description string `db:"description" json:"description"`
	Price       int64  `db:"price" json:"price"`
}

func (u *Offer) String() string {
	return fmt.Sprintf("Offer(%s: %d)", u.ProjectId, u.Price)
}

func (user *Offer) Validate(v *revel.Validation) {

	v.Check(user.Description,
		revel.Required{},
		revel.MaxSize{1500},
	)
}
