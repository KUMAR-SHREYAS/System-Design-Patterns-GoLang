package api

// Maps from api/*.java (DataController, HubController, LoadController)
// Implement HTTP handlers or CLI adapter functions that call into factory, hub, and graph.

import (
	"bus/go_notifyhub/internal/component"
	"bus/go_notifyhub/internal/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DataController struct {
	hubManager *component.HubManager
}

func NewDataController(hm *component.HubManager) *DataController {
	return &DataController{hubManager: hm}
}

// IngestData triggers a publisher to notify observers
func (dc *DataController) IngestData(c *gin.Context) {
	var inputData data.Data
	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic: Find the publisher by sourceID and notify observers
	pub := dc.hubManager.GetPublisher(inputData.Source)
	if pub != nil {
		pub.Notify(&inputData)
		c.JSON(http.StatusOK, gin.H{"status": "Notified observers"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": "Notified observers"})
	}
}
