package controllers

import (
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/revel/revel"
	"going/app/models"
	"going/app/routes"
)

type User struct {
	AbstractController
}

func (c User) Profile() revel.Result {
	user := c.connected()
	if user == nil {
		return c.Redirect(routes.User.LoginPage())
	}
	return c.Render()
}

/**
 * Display the register page
 */
func (c User) Register() revel.Result {
	if user := c.connected(); user != nil {
		return c.Redirect(routes.Projects.List())
	}
	return c.Render()
}

/**
 * Create a new user
 */
func (c User) SaveUser(user models.User, verifyPassword string) revel.Result {
	c.Validation.Required(verifyPassword)
	c.Validation.Required(verifyPassword == user.Password).
		Message("Password does not match")
	user.Validate(c.Validation)

	// validate user data
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.User.Register())
	}
	// hash the password
	user.HashedPassword, _ = bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost)
	err := c.Txn.Insert(&user)
	if err != nil {
		panic(err)
	}
	// save the user to the session
	c.Session["user"] = user.Username
	c.Flash.Success("Welcome, " + user.FirstName)
	// redirect the user
	return c.Redirect(routes.Projects.List())
}

/**
 * Display the login page
 */
func (c User) LoginPage() revel.Result {
	// checks if the user is connected
	if user := c.connected(); user != nil {
		return c.Redirect(routes.Projects.List())
	}
	// display the page
	return c.Render()
}

/**
 * Login the user
 */
func (c User) Login(username, password string, remember bool) revel.Result {
	// get the user from his username
	user := c.getUser(username)
	if user != nil {
		// compare password
		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
		if err == nil {
			// save the user to the session
			c.Session["user"] = username
			if remember {
				c.Session.SetDefaultExpiration()
			} else {
				c.Session.SetNoExpiration()
			}
			c.Flash.Success("Welcome, " + username)
			// redirect the user
			return c.Redirect(routes.Projects.List())
		}
	}

	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(routes.User.LoginPage())
}

/**
 * Logout the user
 */
func (c User) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.Application.Index())
}

/**
 * Display the history of the user
 */
func (c User) History() revel.Result {
	user := c.connected()
	if user == nil {
		return c.Redirect(routes.Application.Index())
	}
	var transactions []models.Transaction
	_, err := c.Txn.Select(&transactions, `select * from Transaction where user_id = ? order by date;`, user.Id)
	if err != nil {
		panic(err)
	}

	var total int64 = 0

	for index, transaction := range transactions {
		var project models.Project
		err := c.Txn.SelectOne(&project, `select * from Project where id = ?`, transaction.ProjectId)
		if err != nil {
			panic(err)
		}
		transaction.Project = project
		transactions[index] = transaction
		total += transaction.Amount
	}

	return c.Render(transactions, total)
}
