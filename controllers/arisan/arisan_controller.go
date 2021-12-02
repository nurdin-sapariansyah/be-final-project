package arisan

import (
	"net/http"
	"strconv"

	"arisan.com/arisan/configs"
	"arisan.com/arisan/models/arisan"
	"arisan.com/arisan/models/response"
	"github.com/labstack/echo/v4"
)

func CreatArisanController(c echo.Context) error {
	var newArisan arisan.Arisan
	c.Bind(&newArisan)

	result := configs.DB.Create(&newArisan)

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
		newArisan,
	}
	return c.JSON(http.StatusOK, response)
}

func GetArisanController(c echo.Context) error {
	var arisan []arisan.Arisan

	result := configs.DB.Preload("Anggota").Find(&arisan)

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
		arisan,
	}
	return c.JSON(http.StatusOK, response)
}

func GetDetailArisanController(c echo.Context) error {
	arisanId, _ := strconv.Atoi(c.Param("arisanId"))

	var data = arisan.Arisan{} // DB

	var response = response.BaseResponse{
		arisanId,
		"success",
		data,
	}
	return c.JSON(http.StatusOK, response)
}