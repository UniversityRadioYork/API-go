package utils

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

var Database *sql.DB

func databaseConnection(c *Config) (*sql.DB) {
	connstr := fmt.Sprintf("host=%s port=%s user=%s "+
    	"password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName)
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