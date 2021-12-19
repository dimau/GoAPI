package services

import (
	"net/http"
	"testing"

	"github.com/Dimau/GoAPI/domain"
	"github.com/Dimau/GoAPI/utils"
	"github.com/stretchr/testify/assert"
)

var (
	userDAOMock usersDAOMock

	// Объявляем функцию, а затем можем ее подменять в каждом конкретном тест-кейсе так,
	// чтобы она возвращала нужное для этого тест-кейса значение
	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDAO = &usersDAOMock{}
}

type usersDAOMock struct{}

func (u *usersDAOMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestUserNotFoundInDB(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "user 0 was not found",
		}
	}
	user, err := UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Message, "user 0 was not found")
}
