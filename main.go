package main

import (
	"io"
	"os"
	"log"
	"net/http"
	"github.com/BurntSushi/toml"
	ut "github.com/UniversityRadioYork/API-go/utils"
)

func main() {
	f, err := os.OpenFile("API-go.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil{
		log.Println(err)
	}
	defer f.Close()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	multi :=  io.MultiWriter(f, os.Stdout)
	log.SetOutput(multi)

	log.Printf("API started")

	config := &ut.Config{}
	_, err = toml.DecodeFile("config.toml", config)
	if err != nil {
		log.Fatal(err)
	}

	ut.NewDatabaseConnection(config)
	ut.NewMemcacheClient(config)

	router := NewRouter()
	portString := ":" + config.ServerPort
	log.Fatal(http.ListenAndServe(portString, router))
}