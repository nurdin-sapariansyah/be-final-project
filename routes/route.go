package routes

import (
	"arisan.com/arisan/controllers/anggota"
	"arisan.com/arisan/controllers/arisan"
	"arisan.com/arisan/controllers/pembayaran"
	"arisan.com/arisan/controllers/undian"
	"arisan.com/arisan/middlewares"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	middlewares.LogMiddleware(e)

	//routes
	//anggota
	e.POST("/register", anggota.RegisterController)
	e.GET("/anggota", anggota.GetAnggotaController)
	e.POST("/login", anggota.LoginController)

	//arisan
	e.GET("/arisan", arisan.GetArisanController)
	e.POST("/newArisan", arisan.CreatArisanController)
	e.POST("/:arisanId", arisan.GetDetailArisanController)

	//undian
	e.GET("/undian", undian.GetUndianController)
	e.POST("/newUndian", undian.CreateUndianController)
	
	e.GET("/pembayaran", pembayaran.GetStatusPembayaranController)

	return e
}
