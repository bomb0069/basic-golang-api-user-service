package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()
	route := e.Group("/api/v1")
	route.GET("/users", getUser)
	e.Logger.Fatal(e.Start(":8080"))
}

func getUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello world")
}
