package user

import (
	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()
	router := e.Group("/api/v1")

	NewUserAPI(router)

	e.Logger.Fatal(e.Start(":8080"))
}
