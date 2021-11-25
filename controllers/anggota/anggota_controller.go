package anggota

import (
	"net/http"

	"arisan.com/arisan/models/anggota"
	"arisan.com/arisan/models/response"
	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	
	var anggota anggota.Anggota
	c.Bind(&anggota)

	if anggota.Password == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			http.StatusBadRequest,
			"Password Kosong",
			nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		http.StatusOK,
		"succes",
		anggota,
	})

}
