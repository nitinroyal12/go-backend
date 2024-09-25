package routes

import (
	"github.com/gorilla/mux"
	"github.com/nitin/cmd/routes/userRoutes"
)

func  MainHandler(router *mux.Router) {
	routes := router.NewRoute()

	// user routes
	user := routes.PathPrefix("/user").Subrouter()
	userRoutes := userroutes.UserHanlder()
	userRoutes.UserHanlder(user) 

}
