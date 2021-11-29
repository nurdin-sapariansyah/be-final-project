package anggota

import (
	"net/http"

	"arisan.com/arisan/configs"
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

func RegisterController(c echo.Context) error {
	var register anggota.Anggota
	c.Bind(&register)

	result := configs.DB.Create(&register)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			http.StatusInternalServerError,
			result.Error.Error(),
			nil,
		})
	}

	var response = response.BaseResponse{
		http.StatusOK,
		"success",
		register,
	}
	return c.JSON(http.StatusOK, response)
}

func GetAnggotaController(c echo.Context) error {
	var anggota []anggota.Anggota

	result := configs.DB.Find(&anggota)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			http.StatusInternalServerError,
			result.Error.Error(),
			nil,
		})
	}

	var response = response.BaseResponse{
		http.StatusOK,
		"success",
		anggota,
	}
	return c.JSON(http.StatusOK, response)
}
