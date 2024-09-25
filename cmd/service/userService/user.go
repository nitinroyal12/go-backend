package userservice

import (
	"context"

	"github.com/nitin/cmd/db"
	"github.com/nitin/cmd/model"
)

func CreateService(user *model.User) (*model.User,error) {
	if _ , err := db.User.InsertOne(context.Background(),user); err != nil {
		return nil ,err
	}
	 
	return user , nil
}