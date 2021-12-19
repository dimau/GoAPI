package domain

import (
	"net/http"
	"testing"
)

var (
	userDaoMock usersDaoMock
)

type usersDaoMock struct{}

func TestGetUserNoUserFound(t *testing.T) {
	// Вызываем тестируемую функцию из этого же пакета
	user, err := UserDAO.GetUser(0)

	// Проверяем, что с такими входными данными функция возвращает nil вместо user
	if user != nil {
		t.Error("we were not expecting a user with id = 0")
	}

	// Проверяем, что с такими входными данными функция возвращает ошибку
	if err == nil {
		t.Error("we were expecting an error when user id is 0")
	}

	// Проверяем корректность статуса в описании ошибки
	if err.StatusCode != http.StatusNotFound {
		t.Error("we were expecting 404 when user is not found")
	}
}
