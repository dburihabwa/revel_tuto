package tests

import (
	"github.com/revel/revel"
	"net/url"
)

type UserTest struct {
	revel.TestSuite
}

func (t *UserTest) Before() {
	t.PostForm("/register", url.Values{
		"user.Username":  {"test"},
		"user.Password":  {"testtest"},
		"user.FirstName": {"test"},
		"user.LastName":  {"test"},
		"verifyPassword": {"testtest"},
	})
	t.Get("/logout")
}

func (t *UserTest) TestCreateUser() {
	t.PostForm("/register", url.Values{
		"user.Username":  {"test"},
		"user.Password":  {"testtest"},
		"user.FirstName": {"test"},
		"user.LastName":  {"test"},
		"verifyPassword": {"testtest"},
	})
	t.AssertOk()
	t.AssertStatus(200)
	t.AssertEqual("/projects", t.Response.Request.URL.Path)
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *UserTest) TestLogin() {
	t.PostForm("/login", url.Values{
		"username": {"test"},
		"password": {"testtest"},
		"remember": {"false"},
	})
	t.AssertStatus(200)
	t.AssertEqual("/projects", t.Response.Request.URL.Path)
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *UserTest) TestLoginFail() {
	t.PostForm("/login", url.Values{
		"username": {"bad"},
		"password": {"bad"},
		"remember": {"false"},
	})
	t.AssertStatus(200)
	t.AssertEqual("/login", t.Response.Request.URL.Path)
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *UserTest) TestProfileNotConnected() {
	t.Get("/profile")
	t.AssertStatus(200)
	t.AssertEqual("/login", t.Response.Request.URL.Path)
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *UserTest) TestProfile() {
	t.PostForm("/login", url.Values{
		"username": {"test"},
		"password": {"testtest"},
		"remember": {"false"},
	})
	t.Get("/profile")
	t.AssertStatus(200)
	t.AssertEqual("/profile", t.Response.Request.URL.Path)
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *UserTest) After() {
	t.Get("/logout")
	t.Delete("/user/test")
}
