package user

import (
	"net/http"

	"github.com/nitin/cmd/model"
	userservice "github.com/nitin/cmd/service/userService"
	"github.com/nitin/cmd/utils"
	"github.com/nitin/cmd/validation"
)

func Create(w http.ResponseWriter,r *http.Request) {
	var user *model.User
	if err := utils.Prasejson(r,&user); err != nil {
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}
	if err := validation.CheckValidation(user); err != nil {
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}
	result, err := userservice.CreateService(user)
	if err != nil {
		utils.WriteError(w,http.StatusNotFound,err)
		return
	}
	utils.WriteJson(w,http.StatusOK,result)
}