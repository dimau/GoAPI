package services

import (
	"github.com/Dimau/GoAPI/domain"
	"github.com/Dimau/GoAPI/utils"
)

var (
	UsersService usersService
)

type usersService struct{}

func (u *usersService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDAO.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
