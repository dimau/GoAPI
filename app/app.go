package app

import (
	"net/http"

	"github.com/Dimau/GoAPI/controllers"
)

func SrartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
