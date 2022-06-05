package events

import (
	"net/http"
	"time"
	"tusa/internal/model"
	"tusa/internal/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TusaEventHandler interface {
	List(c *gin.Context)
	Add(c *gin.Context)
}

type tusaEventHandler struct {
	storage storage.TusaEventStore
}

func NewEventHandler(storage storage.TusaEventStore) TusaEventHandler {
	return &tusaEventHandler{
		storage: storage,
	}
}

func (h tusaEventHandler) List(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"events": h.storage.FindAllLatest(time.Now())})
}
func (h *tusaEventHandler) Add(c *gin.Context) {
	var event model.TusaEvent
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Failed to Parse": err.Error()})
		return
	}
	event.Id = uuid.New()
	h.storage.Add(event)
	c.JSON(http.StatusOK, gin.H{"event": event})
}
