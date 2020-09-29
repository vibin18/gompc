package services

import "github.com/vibin18/gompc/mvc/domain"

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)

}
