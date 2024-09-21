package routes

import (
	"github.com/gorilla/mux"
	todoroutes "github.com/nitin/cmd/routes/todoRoutes"
	"github.com/nitin/cmd/routes/userRoutes"
)

func  MainHandler(router *mux.Router) {
	routes := router.NewRoute()

	// user routes
	user := routes.PathPrefix("/user").Subrouter()
	userRoutes := userroutes.UserHanlder()
	userRoutes.UserHanlder(user) 
	
	// todo route
	todo := router.PathPrefix("/todo").Subrouter()
	todoRoute := todoroutes.TodoHandler()
	todoRoute.TodoHanlder(todo)
	

}
