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
	portString := ":" + config.ServerPort
	log.Fatal(http.ListenAndServe(portString, router))
}