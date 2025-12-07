package main

import (
	"github.com/gin-gonic/gin"

	"bus/go_notifyhub/internal/api"
	"bus/go_notifyhub/internal/component"
	"bus/go_notifyhub/internal/service"
)

func main() {
	// 1. Core Services
	graph := service.NewGraph()
	hubManager := component.NewHubManager(graph)

	// 2. Controllers (Injecting Manager)
	hubController := api.NewHubController(hubManager) // From previous step
	dataController := api.NewDataController(hubManager)
	loadController := api.NewLoadController(hubManager)

	// 3. Gin Router
	r := gin.Default()

	// 4. Routes
	apiGroup := r.Group("/api")
	{
		// Hub Routes
		hubController.RegisterRoutes(r) // Assuming you use the logic from previous response

		// Data Routes
		apiGroup.POST("/data", dataController.IngestData)

		// Load Routes
		apiGroup.GET("/load-init", loadController.InitLoad)
	}

	// 5. Start
	r.Run(":8080")
}
