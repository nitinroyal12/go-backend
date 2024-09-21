package todoroutes

import (
	"github.com/gorilla/mux"
	"github.com/nitin/cmd/controller/todo"
)

type Handler struct {
}

func TodoHandler() *Handler {
	return &Handler{}
}

func (h *Handler) TodoHanlder(router *mux.Router) {
	router.HandleFunc("/create", todo.Create)
}
