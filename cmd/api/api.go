package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nitin/cmd/config"
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
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()
	routes.MainHandler(subRouter)
	log.Println("server is running on port" + config.Keys().Port )
	return  http.ListenAndServe( a.addr,router); 
}
