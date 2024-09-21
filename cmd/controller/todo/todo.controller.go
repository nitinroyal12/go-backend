package todo

import "net/http"

func Create(w http.ResponseWriter , r *http.Request) {
	w.Write([]byte("todo route is also working"))
}