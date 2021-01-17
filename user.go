package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewUserAPI(group *echo.Group, resource *Resource) {
	group.GET("/users", getAllUsersHandler(resource))
}

func getAllUsersHandler(resource *Resource) func(echo.Context) error {
	return func(ctx echo.Context) error {
		repository := NewUserRepository(resource)

		u, _ := repository.GetAll()

		return ctx.JSON(http.StatusOK, u)
	}
}
