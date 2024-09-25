package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Prasejson(r *http.Request,payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJson(w http.ResponseWriter,status int , v any) error  {
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError (w http.ResponseWriter,status int ,err error){
	WriteJson(w,status,map[string]string{"error":err.Error()})
}