package model

type User struct {
	FullName string `bson:"fullname,omitempty"`
	UserName string `bson:"username,omitempty"`
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}


