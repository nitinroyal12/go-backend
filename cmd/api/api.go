package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nitin/cmd/config"
	"github.com/nitin/cmd/db"
	"github.com/nitin/cmd/routes"
)

type api struct {
	addr string
}

func InitServer(addr string) *api {
	return &api{
		addr: addr,
	}
}

func (a *api) Start() error {
	
	// db connection
	if err := db.InitDatabase(config.Keys().DB); err != nil {
		log.Fatal(err)
	}

	// routing setup and server
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()
	routes.MainHandler(subRouter)
	log.Println("server is running on port" + config.Keys().Port )
	return  http.ListenAndServe( a.addr,router); 


}
