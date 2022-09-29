package city

import (
	"crud-api/dto/api"
	"crud-api/dto/city"
	"crud-api/entity"
	"crud-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strings"
)

func CreateCity(c *gin.Context) {
	var input city.RequestCity
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, api.Response{
			OK:       false,
			Response: nil,
			Errors: []api.Error{
				{
					Code:    api.ErrorCode400InvalidData,
					Message: api.ErrorMessage400InvalidData,
				},
			},
		})
		return
	}
	input.Code = strings.ToLower(input.Code)

	db := services.GetOrmService()
	result := db.Where(&entity.City{Code: input.Code}).First(&entity.City{})

	if result.Error == nil {
		c.JSON(400, api.Response{
			OK:       false,
			Response: nil,
			Errors: []api.Error{{
				Code:    api.ErrorCode400CityExists,
				Message: api.ErrorMessage400CityExists,
			}},
		})
		return
	}
	match, _ := regexp.MatchString("^[\u0600-\u06FF\\s |\\s]+$", input.Name)
	if !match {

		c.JSON(400, api.Response{
			OK:       false,
			Response: nil,
			Errors: []api.Error{{
				Code:    api.ErrorCode400CityNameEnglishOnly,
				Message: api.ErrorMessage400CityNameEnglishOnly,
			}},
		})
		return
	}

	// Create new city
	newCity := entity.City{Name: input.Name, Code: input.Code}
	db.Create(&newCity)

	c.JSON(http.StatusOK, api.Response{
		OK: true,
		Response: city.ResponseCity{
			ID:   newCity.ID,
			Name: newCity.Name,
			Code: newCity.Code,
		},
		Errors: nil,
	})
}
