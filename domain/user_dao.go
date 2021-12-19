package domain

import (
	"fmt"
	"net/http"

	"github.com/Dimau/GoAPI/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Dmitrii", LastName: "Ushakov", Email: "dimau@dfdf.ru"},
	}

	// Объявляем публичную переменную UserDAO для работы с пользователями в БД
	// И при объявлении говорим, что эта переменная может содержать значения любых типов, поддерживающих интерфейс usersDAOInterface
	UserDAO usersDAOInterface
)

func init() {
	// При инициализации записываем в нашу публичную переменную значение с типом userDAO
	// Это позволит сервисам при обращении к этой переменной вызывать метод GetUser объявленный для данного типа
	UserDAO = &userDAO{}
}

type usersDAOInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDAO struct{}

func (u *userDAO) GetUser(userId int64) (*User, *utils.ApplicationError) {
	user := users[userId]
	if user == nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("user %v was not found", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}

	return user, nil
}
