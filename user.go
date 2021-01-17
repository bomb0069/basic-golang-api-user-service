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
		u := Users{
			User{
				Firstname: "f1",
				Lastname:  "l1",
				Title:     "Mr.",
			},
			User{
				Firstname: "f2",
				Lastname:  "l2",
				Title:     "Miss.",
			},
		}
		return ctx.JSON(http.StatusOK, u)
	}
}

func getUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello world")
}
