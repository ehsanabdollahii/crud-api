package city

import (
	"crud-api/dto/api"
	"crud-api/entity"
	"crud-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Delete a CityResponse
func DeleteCity(c *gin.Context) {
	// Get city if exist
	var cityEntity entity.City
	db := services.GetOrmService()

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

	db.Unscoped().Delete(&cityEntity)

	c.JSON(http.StatusNoContent, nil)
}
