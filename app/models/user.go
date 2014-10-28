package models

import (
	"fmt"
	"github.com/revel/revel"
	"regexp"
)

type User struct {
	Id             int64  `db:"id" json:"id"`
	FirstName      string `db:"firstname" json:"firstname"`
	LastName       string `db:"lastname" json:"lastname"`
	Username       string `db:"username" json:"username"`
	Password       string `db:"-"`
	HashedPassword []byte `db:"password" json:"-"`
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s: %s)", u.Id, u.Username)
}

var userRegex = regexp.MustCompile("^\\w*$")

func (user *User) Validate(v *revel.Validation) {
	v.Check(user.Username,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{4},
		revel.Match{userRegex},
	)

	ValidatePassword(v, user.Password).
		Key("user.Password")

	v.Check(user.FirstName,
		revel.Required{},
		revel.MaxSize{100},
	)
	v.Check(user.LastName,
		revel.Required{},
		revel.MaxSize{100},
	)
}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	)
}
