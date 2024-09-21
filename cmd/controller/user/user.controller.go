package user

import "net/http"

func Create(w http.ResponseWriter,r *http.Request) {
	w.Write([]byte("server is perfect"))
}