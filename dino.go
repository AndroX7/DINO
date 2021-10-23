package main

import (
	"DINO/dynowebportal"
	"encoding/json"
	"log"
	"os"
)

//configuration is
type configuration struct {
	Webserver string `json:"webserver"`
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	config := new(configuration)
	json.NewDecoder(file).Decode(&config)

	log.Println("Starting web server on address: ", config.Webserver)
	dynowebportal.RunWebPortal(config.Webserver)
}
