package routes

import (
	"arisan.com/arisan/constants"
	"arisan.com/arisan/controllers/anggota"
	"arisan.com/arisan/controllers/arisan"
	"arisan.com/arisan/controllers/pembayaran"
	"arisan.com/arisan/controllers/undian"
	"arisan.com/arisan/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	middlewares.LogMiddleware(e)

	config := middleware.JWTConfig{
		Claims:     &middlewares.JwtCustomClaims{},
		SigningKey: []byte(constants.SECRET_JWT),
	}
	
	//routes
	//anggota
	eAnggota := e.Group("/anggota")
	eAnggota.Use(middleware.JWTWithConfig(config))
	eAnggota.GET("/", anggota.GetAnggotaController)
	e.POST("/register", anggota.RegisterController)
	e.POST("/login", anggota.LoginController)

	//arisan
	eArisan := e.Group("/arisan")
	eArisan.Use(middleware.JWTWithConfig(config))
	eArisan.GET("/", arisan.GetArisanController)
	eArisan.POST("/", arisan.CreatArisanController)
	eArisan.GET("/:arisanId", arisan.GetDetailArisanController)

	//undian
	eUndian := e.Group("/undian")
	eUndian.GET("/", undian.GetUndianController)
	eUndian.POST("/", undian.CreateUndianController)
	eUndian.Use(middleware.JWTWithConfig(config))
	
	e.GET("/pembayaran", pembayaran.GetStatusPembayaranController)

	return e
}
