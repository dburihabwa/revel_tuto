package models

import (
	"fmt"
	"github.com/revel/revel"
	"time"
)

type Transaction struct {
	Id        int64     `db:"id" json:"id"`
	UserId    int64     `db:"user_id" json:"user_id"`
	User      User      `db:"-"`
	ProjectId int64     `db:"project_id" json:"project_id"`
	Project   Project   `db:"-"`
	date      time.Time `db:"date" json:"date"`
}

func (u *Transaction) String() string {
	return fmt.Sprintf("Transaction(%s: %s)", u.Id)
}

func (transaction *Transaction) Validate(v *revel.Validation) {

}
