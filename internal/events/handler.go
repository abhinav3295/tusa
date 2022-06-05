package events

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TusaEventHandler interface {
	List(c *gin.Context)
	Add(c *gin.Context)
}

type tusaEventHandler struct {
	events []TusaEvent
}

func NewEventHandler() TusaEventHandler {
	return &tusaEventHandler{
		events: make([]TusaEvent, 0),
	}
}

func (h tusaEventHandler) List(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"events": h.events})
}
func (h *tusaEventHandler) Add(c *gin.Context) {
	var event TusaEvent
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Failed to Parse": err.Error()})
		return
	}
	h.events = append(h.events, event)
	c.JSON(http.StatusOK, gin.H{"event": event})
}
