package routes

import (
	"arisan.com/arisan/controllers/anggota"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/login", anggota.LoginController)

	return e
}
