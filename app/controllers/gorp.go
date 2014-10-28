package controllers

import (
	"database/sql"
	"fmt"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
	"going/app/models"
	"strings"
)

var (
	Dbm *gorp.DbMap
)

func getParamString(param string, defaultValue string) string {
	p, found := revel.Config.String(param)
	if !found {
		if defaultValue == "" {
			revel.ERROR.Fatal("Cound not find parameter: " + param)
		} else {
			return defaultValue
		}
	}
	return p
}

func getConnectionString() string {
	host := getParamString("db.host", "localhost")
	port := getParamString("db.port", "3306")
	user := getParamString("db.user", "root")
	pass := getParamString("db.password", "root")
	dbname := getParamString("db.name", "going")
	protocol := getParamString("db.protocol", "tcp")
	dbargs := getParamString("dbargs", " ")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		user, pass, protocol, host, port, dbname, dbargs)
}

func InitDB() {
	connectionString := getConnectionString()
	if db, err := sql.Open("mysql", connectionString); err != nil {
		revel.ERROR.Fatal(err)
	} else {
		Dbm = &gorp.DbMap{
			Db:      db,
			Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	}
	// Defines the table for use by GORP
	// This is a function we will create soon.
	defineTable(Dbm)
	if err := Dbm.CreateTablesIfNotExists(); err != nil {
		revel.ERROR.Fatal(err)
	}
}

func defineTable(dbm *gorp.DbMap) {
	// set "id" as primary key and autoincrement
	user := dbm.AddTable(models.User{}).SetKeys(true, "id")
	user.ColMap("username").SetMaxSize(25)

	project := dbm.AddTable(models.Project{}).SetKeys(true, "id")
	project.ColMap("description").SetMaxSize(1500)

	offer := dbm.AddTable(models.Offer{}).SetKeys(true, "id")
	offer.ColMap("description").SetMaxSize(1500)

	//transaction := dbm.AddTable(models.Transaction{}).SetKeys(true, "id")
}

type GorpController struct {
	*revel.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() revel.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
