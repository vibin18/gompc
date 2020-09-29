package domain

import (
	"errors"
	"fmt"
)

var (
	users = map[int64] *User{
		777: {Id: 777, FirstName: "Vibin", LastName: "Manavalan", email: "vibin@vibin.com"},
	}
)
func GetUser(userId int64) (*User, error){
	user := users[userId]
	if user == nil {
		return nil,errors.New(fmt.Sprintf("User %v not found", userId))
	}
	return user, nil
}