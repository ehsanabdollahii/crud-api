package city

import (
	"crud-api/dto/api"
	"crud-api/dto/city"
	"crud-api/entity"
	"crud-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
)

func UpdateCity(c *gin.Context) {
	// Get model if exist
	var cityEntity entity.City
	db := services.GetOrmService()

	var input city.RequestUpdateCity
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

	cityID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
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

	if err := db.Where(&entity.City{ID: uint(cityID)}).First(&cityEntity).Error; err != nil {
		c.JSON(http.StatusBadRequest, api.Response{
			OK:       false,
			Response: nil,
			Errors: []api.Error{
				{
					Code:    api.ErrorCode400CityNotExists,
					Message: api.ErrorMessage400CityNotExists,
				},
			},
		})
		return
	}
	if input.Name != cityEntity.Name {
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
		cityEntity.Name = input.Name
	}

	db.Save(cityEntity)

	c.JSON(http.StatusOK, api.Response{
		OK: true,
		Response: city.ResponseCity{
			ID:   cityEntity.ID,
			Name: cityEntity.Name,
			Code: cityEntity.Code,
		},
		Errors: nil,
	})
}
