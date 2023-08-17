package routes

import (
	"github.com/Asad2730/EventManagement/service"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	r.GET("/", service.GetUsers)
	r.GET("/:id", service.GetUser)
	r.POST("/", service.CreateUser)
	r.PUT("/:id", service.UpdateUser)
	r.DELETE("/:id", service.DeleteUser)
}
