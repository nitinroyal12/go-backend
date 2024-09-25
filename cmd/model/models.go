package model

type User struct {
	FullName string `bson:"fullname,omitempty" validate:"required,min=2,max=32"`
	UserName string `bson:"username,omitempty,unique" validate:"required,min=2,max=32"`
	Email    string `bson:"email,omitempty,unique" validate:"required,email"`
	Password string `bson:"password,omitempty" validate:"required,min=6,max=32"`
}
