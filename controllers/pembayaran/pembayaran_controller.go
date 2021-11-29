package pembayaran

import (
	"net/http"

	"arisan.com/arisan/configs"
	"arisan.com/arisan/models/anggota"
	"arisan.com/arisan/models/response"
	"github.com/labstack/echo/v4"
)

func GetStatusPembayaranController(c echo.Context) error {
	var statusPembayaran []anggota.Anggota

	result := configs.DB.Preload("Status").Find(&statusPembayaran)

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
		statusPembayaran,
	}
	return c.JSON(http.StatusOK, response)
}