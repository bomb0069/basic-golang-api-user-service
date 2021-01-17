package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewUserAPI(group *echo.Group) {
	group.GET("/users", getUser)
}

func getUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello world")
}
