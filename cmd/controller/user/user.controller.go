package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nitin/cmd/model"
	userservice "github.com/nitin/cmd/service/userService"
	"github.com/nitin/cmd/utils"
)

func Create(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	var user model.User
	if err := utils.Prasejson(r,user); err != nil {
		utils.WriteError(w,http.StatusBadRequest,err)
	}
	result, err := userservice.CreateService(user)
	if err != nil {
		json.NewEncoder(w).Encode("user not created somthing went wrong")
		log.Fatal(err)
		return
	}
	json.NewEncoder(w).Encode(result)
}