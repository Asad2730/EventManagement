package main

import (
	"github.com/Asad2730/EventManagement/connect"
	"github.com/Asad2730/EventManagement/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	connect.ConnectAws()
}

func main() {

	r := gin.Default()

	eventGroup := r.Group("/event")
	userGroup := r.Group("/user")

	routes.EventsRoutes(eventGroup)
	routes.UserRoutes(userGroup)
	r.Run()

}
