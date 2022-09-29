package clients

import (
	"crud-api/model/city"
	"crud-api/view/city_view"
	"github.com/gin-gonic/gin"
)

func ApplyClientRouter(router *gin.Engine) {
	versionOne := router.Group("/api/v1")
	{

		cityGroup := versionOne.Group("/cities")
		{
			cityGroup.POST("/", city.CreateCity)
			cityGroup.GET("/:id", city_view.FindCityByID)
			cityGroup.GET("/", city_view.FindCity)
			cityGroup.PATCH("/:id", city.UpdateCity)
			cityGroup.DELETE("/:id", city.DeleteCity)
		}
	}
}
