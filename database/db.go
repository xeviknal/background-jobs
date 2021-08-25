package database

import (
	"database/sql"
	"gopkg.in/gorp.v1"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var dbmap *gorp.DbMap

// Singleton for access to the DB config
func GetDb() *gorp.DbMap {
	if dbmap != nil {
		return dbmap
	}
	dbmap = NewDatabase()
	return dbmap
}

func NewDatabase() *gorp.DbMap {
	db, err := sql.Open("mysql", "jobs:jobs@tcp(localhost)/jobs?parseTime=true")
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
	if err := CreateScheme(); err != nil {
		log.Fatalf("Error creating the scheme and tables: %v", err)
		return dbmap
	}
	return dbmap
}
