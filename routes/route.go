package routes

import (
	"arisan.com/arisan/controllers/anggota"
	"arisan.com/arisan/controllers/arisan"
	"arisan.com/arisan/controllers/pembayaran"
	"arisan.com/arisan/controllers/undian"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	//routes
	//anggota
	e.POST("/register", anggota.RegisterController)
	e.GET("/anggota", anggota.GetAnggotaController)
	e.POST("/login", anggota.LoginController)

	//arisan
	e.GET("/arisan", arisan.GetArisanController)
	e.POST("/new", arisan.CreatArisanController)

	//undian
	e.GET("/undian", undian.GetUndianController)
	e.POST("/create", undian.CreUndianController)
	
	e.GET("/pembayaran", pembayaran.GetStatusPembayaranController)

	return e
}
