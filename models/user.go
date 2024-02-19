package models

import (
	"github.com/TvGelderen/algo-alcove/database"
)

type User struct {
	Email    string
	Username string
	Role     string
}

func ToUser(dbModel database.User) User {
	return User{
		Email:    dbModel.Email,
		Username: dbModel.Username,
		Role:     dbModel.Role,
	}
}
