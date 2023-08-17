package routes

import (
	"github.com/Asad2730/EventManagement/service"
	"github.com/gin-gonic/gin"
)

func EventsRoutes(r *gin.RouterGroup) {
	r.GET("/", service.GetEvents)
	r.GET("/:id", service.GetEvent)
	r.POST("/", service.CreateEvent)
	r.PUT("/:id", service.UpdateEvent)
	r.DELETE("/:id", service.DeleteEvent)
}
