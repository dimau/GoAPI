package services

import (
	"github.com/Dimau/GoAPI/domain"
	"github.com/Dimau/GoAPI/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
