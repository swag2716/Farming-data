package routes

import (
	"Farming_data/controllers"

	"github.com/gin-gonic/gin"
)

func FarmingRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("/create_country", controllers.CreateCountry())
	incomingRoutes.POST("/create_farmer", controllers.CreateFarmer())
	incomingRoutes.POST("/create_farm", controllers.CreateFarm())
	incomingRoutes.POST("/create_schedule", controllers.CreateSchedule())
	incomingRoutes.GET("/get_due_schedules", controllers.GetDueSchedules())
}
