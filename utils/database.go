package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var Database *sql.DB

func databaseConnection(c *Config) *sql.DB {
	connstr := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Database connected")
	}
	return db
}

func NewDatabaseConnection(c *Config) {
	Database = databaseConnection(c)
}
