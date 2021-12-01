package undian

import (
	"net/http"

	"arisan.com/arisan/configs"
	"arisan.com/arisan/models/response"
	"arisan.com/arisan/models/undian"
	"github.com/labstack/echo/v4"
)

func CreUndianController(c echo.Context) error {
	var newUndian undian.Undian
	c.Bind(&newUndian)

	result := configs.DB.Create(&newUndian)

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
		newUndian,
	}
	return c.JSON(http.StatusOK, response)
}

func GetUndianController(c echo.Context) error {
	var undian []undian.Undian

	result := configs.DB.Find(&undian)

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
		undian,
	}
	return c.JSON(http.StatusOK, response)
}