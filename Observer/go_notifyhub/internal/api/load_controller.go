package api

import (
	"bus/go_notifyhub/internal/component"
	"bus/go_notifyhub/internal/core"

	"github.com/gin-gonic/gin"
)

type LoadController struct {
	hubManager *component.HubManager
}

func NewLoadController(hm *component.HubManager) *LoadController {
	return &LoadController{hubManager: hm}
}

// InitLoad creates dummy publishers/observers for testing
func (lc *LoadController) InitLoad(c *gin.Context) {
	p1 := core.NewPublisher("pub-1")
	lc.hubManager.RegisterPublisher(p1)

	o1 := core.NewObserver("obs-1")
	lc.hubManager.RegisterObserver(o1)

	// 2. Link them using the Manager (Updates Graph + Publisher)
	err := lc.hubManager.Link("pub-1", "obs-1")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Load initialized: pub-1 -> obs-1"})
}
