package city_view

import (
	"crud-api/dto/api"
	"crud-api/dto/city"
	"crud-api/entity"
	"crud-api/services"
	"github.com/gin-gonic/gin"
)

func FindCity(c *gin.Context) {
	var cityEntities []*entity.City

	db := services.GetOrmService()

	db.Find(&cityEntities)

	cities := make([]*city.ResponseCity, len(cityEntities))

	for index, cityEntity := range cityEntities {
		cities[index] = &city.ResponseCity{
			ID:   cityEntity.ID,
			Name: cityEntity.Name,
			Code: cityEntity.Code,
		}
	}
	c.JSON(200, api.Response{
		OK:       true,
		Response: cities,
		Errors:   nil,
	})
}
