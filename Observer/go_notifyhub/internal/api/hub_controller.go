package api

import (
	"bus/go_notifyhub/internal/component"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HubController holds the dependencies required for the handlers.
// This replaces the @Autowired private HubManager field.
type HubController struct {
	hubManager *component.HubManager
}

// NewHubController is a constructor to inject the dependency manually.
func NewHubController(hm *component.HubManager) *HubController {
	return &HubController{
		hubManager: hm,
	}
}

// RegisterRoutes defines the @RequestMapping("/hub") logic.
func (h *HubController) RegisterRoutes(r *gin.Engine) {
	//Create a route group for "/hub"
	hubGroup := r.Group("/hub")
	{
		hubGroup.GET("/all", h.ShowSyncHub)
		hubGroup.GET("/graph", h.ShowGraph)
		hubGroup.GET("/:type", h.ShowHub) // :type captures the path variable
	}
}

// ShowSyncHub corresponds to @GetMapping("/all")
func (h *HubController) ShowSyncHub(c *gin.Context) {
	//http.StatusAccepted is 202 (ResponseEntity.accepted())
	c.JSON(http.StatusAccepted, h.hubManager.GetRegistry())
}

// ShowSyncHub corresponds to @GetMapping("/graph")
func (h *HubController) ShowGraph(c *gin.Context) {
	c.JSON(http.StatusAccepted, h.hubManager.GetG())
}

// ShowHub corresponds to @GetMapping("/{type}")
func (h *HubController) ShowHub(c *gin.Context) {
	//Retrieve the @PathVariable "type"
	t := c.Param("type")
	c.JSON(http.StatusAccepted, h.hubManager.Get(t))
}
