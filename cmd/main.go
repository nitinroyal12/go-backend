package main

import (
	"log"

	"github.com/nitin/cmd/api"
	"github.com/nitin/cmd/config"
)

func main() {
	server := api.InitServer(config.Keys().Port)
	if err := server.Start() ; err != nil {
		log.Fatal(err)
	}
}