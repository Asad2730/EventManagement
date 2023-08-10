package routes

import (
	"github.com/Asad2730/EventManagement/controller"
	"github.com/gin-gonic/gin"
)

func EventsRoutes(r *gin.Engine) {
	r.GET("/", controller.GetAll)
	r.GET("/:id", controller.GetById)
	r.POST("/", controller.Create)
	r.PUT("/:id", controller.Update)
	r.DELETE("/:id", controller.Delete)
}
