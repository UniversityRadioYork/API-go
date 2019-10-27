package main

import (
	"os"
	"log"
	"net/http"
	"github.com/BurntSushi/toml"
	ut "github.com/UniversityRadioYork/API-go/utils"
)

func main() {
	f, err := os.OpenFile("APIgo.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil{
		log.Println(err)
	}
	defer f.Close()
	log.SetFlags(log.Llongfile)
	log.SetOutput(f)

	log.Printf("API started")

	config := &ut.Config{}
	_, err = toml.DecodeFile("config.toml", config)
	if err != nil {
		log.Fatal(err)
	}

	ut.NewDatabaseConnection(config)
	ut.NewMemcacheClient(config)

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}