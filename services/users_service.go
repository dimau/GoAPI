package services

import "github.com/Dimau/GoAPI/domain"

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
