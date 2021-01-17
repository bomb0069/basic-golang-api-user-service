package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewUserAPI(group *echo.Group) {
	group.GET("/users", getAllUsersHandler())
}

func getAllUsersHandler() func(echo.Context) error {
	return func(ctx echo.Context) error {
		repository := NewUserRepository()

		u, _ := repository.GetAll()

		return ctx.JSON(http.StatusOK, u)
	}
}
