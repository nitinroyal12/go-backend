package userroutes

import (
	"github.com/gorilla/mux"
	"github.com/nitin/cmd/controller/user"
)

type Handler struct {
}

func UserHanlder() *Handler {
	return &Handler{}
}

func (h *Handler) UserHanlder(router *mux.Router) {
	router.HandleFunc("/create", user.Create).Methods("POST")
}


