package routes

import (
	"arisan.com/arisan/controllers/anggota"
	"arisan.com/arisan/controllers/pembayaran"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/register", anggota.RegisterController)
	e.GET("/anggota", anggota.GetAnggotaController)
	e.POST("/login", anggota.LoginController)
	e.GET("/pembayaran", pembayaran.GetStatusPembayaranController)

	return e
}
