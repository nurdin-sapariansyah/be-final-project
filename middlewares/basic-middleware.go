package middlewares

import (
	"arisan.com/arisan/configs"
	"arisan.com/arisan/models/anggota"
	"github.com/labstack/echo/v4"
)

func BasicAuth(nomorHp, password string, c echo.Context) (bool, error) {
	var db = configs.DB
	var user anggota.Anggota

	if err := db.Where("email = ? AND password = ?", nomorHp, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil

}
