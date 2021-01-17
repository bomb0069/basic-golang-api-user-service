package user

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func StartServer() {
	e := echo.New()
	router := e.Group("/api/v1")

	resource, err := CreateResource()
	if err != nil {
		logrus.Error(err)
	}
	defer resource.Close()

	NewUserAPI(router, resource)

	e.Logger.Fatal(e.Start(":8080"))
}
