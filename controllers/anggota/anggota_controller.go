package anggota

import (
	"net/http"

	"arisan.com/arisan/configs"
	"arisan.com/arisan/middlewares"
	"arisan.com/arisan/models/anggota"
	"arisan.com/arisan/models/response"
	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {

	var anggota anggota.Anggota
	c.Bind(&anggota)

	if err := configs.DB.Where("nomor_hp = ? AND password = ?", anggota.NomorHp, anggota.Password).First(&anggota).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, response.BaseResponse{
			http.StatusUnauthorized,
			"Nomor HP dan Password tidak sesuai",
			nil,
		})
	}

	if anggota.Password == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			http.StatusBadRequest,
			"Password Kosong",
			nil,
		})
	}

	// generate token
	token, err := middlewares.GenereteTokenJWT(int(anggota.Id))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			http.StatusInternalServerError,
			"Error ketika generate token JWT",
			nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		http.StatusOK,
		"succes",
		map[string]string{
			"token": token,
		},
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
