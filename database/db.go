package database

import (
	"database/sql"
	"fmt"
	"gopkg.in/gorp.v1"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var dbmap *gorp.DbMap
var user, pwd, db string

func SetConnectionConfig(u, p, d string) {
	user = u
	pwd = p
	db = d
}

// Singleton for access to the DB config
func GetDb() *gorp.DbMap {
	if dbmap != nil {
		return dbmap
	}
	dbmap = NewDatabase()
	return dbmap
}

func NewDatabase() *gorp.DbMap {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost)/%s?parseTime=true", user, pwd, db))
	if err != nil {
		log.Fatal(err)
		return dbmap
	}

	// Needs to be tuned according the available infrastructure
	// Allowing 10 connection maximum
	// Having a pool of 3 ready to
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(3)
	db.SetConnMaxLifetime(5 * time.Minute)

	// construct a gorp DbMap
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "jobs:", log.Lmicroseconds))

	// Register and create the tables
	// This should go outside the startup process. Not to mess with databases.
	if err := CreateScheme(); err != nil {
		log.Fatalf("Error creating the scheme and tables: %v", err)
		return dbmap
	}

	return dbmap
}

func Clean() {
	DropTables()
	dbmap = nil
}
