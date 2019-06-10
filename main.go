package main

import (
	"log"
	"net/http"
	"github.com/BurntSushi/toml"
	ut "github.com/UniversityRadioYork/API-go/utils"
)

func main() {
	log.Printf("API started")

	config := &ut.Config{}
	_, err := toml.DecodeFile("config.toml", config)
	if err != nil {
		log.Fatal(err)
	}

	ut.NewDatabaseConnection(config)
	ut.NewMemcacheClient(config)

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}