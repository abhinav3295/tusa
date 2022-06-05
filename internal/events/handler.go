package events

import (
	"fmt"
	"log"
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
	Get(c *gin.Context)
	Rsvp(c *gin.Context)
}

type tusaEventHandler struct {
	tusaStorage storage.TusaEventStore
	rsvpStorage storage.RsvpStore
}

func NewEventHandler(tusaStorage storage.TusaEventStore, rsvpStorage storage.RsvpStore) TusaEventHandler {
	return &tusaEventHandler{
		tusaStorage: tusaStorage,
		rsvpStorage: rsvpStorage,
	}
}

func (h tusaEventHandler) Get(c *gin.Context) {
	event, exist := h.find(c)
	if exist {
		c.JSON(http.StatusOK, gin.H{"event": event})
		return
	}
	eventIdStr := c.Param("id")
	c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Event with id '%s' does not exist", eventIdStr)})
}

func (h tusaEventHandler) find(c *gin.Context) (model.TusaEvent, bool) {
	eventIdStr := c.Param("id")
	if eventId, err := uuid.Parse(eventIdStr); err == nil {
		return h.tusaStorage.Find(eventId)
	} else {
		log.Panicln(err.Error())
		return model.TusaEvent{}, false
	}
}

func (h tusaEventHandler) List(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"events": h.tusaStorage.FindAllLatest(time.Now())})
}
func (h tusaEventHandler) Add(c *gin.Context) {
	var event model.TusaEvent
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Failed to Parse": err.Error()})
		return
	}
	event.Id = uuid.New()
	h.tusaStorage.Add(event)
	c.JSON(http.StatusOK, gin.H{"event": event})
}

func (h tusaEventHandler) Rsvp(c *gin.Context) {
	event, exist := h.find(c)
	if exist {
		var rsvp model.RSVP
		if err := c.ShouldBindJSON(&rsvp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Failed to Parse": err.Error()})
			return
		}
		rsvp.Id = uuid.New()
		rsvp.EventId = event.Id
		// if userId, err := uuid.Parse(c.GetString("AuthenticatedUser")); err != nil {
		// 	c.JSON(http.StatusForbidden, gin.H{"error": "User is not authenticated"})
		// 	return
		// } else {
		// 	rsvp.UserId = userId
		// }
		h.rsvpStorage.AddOrUpdate(rsvp)
		c.JSON(http.StatusOK, gin.H{"rsvp": rsvp})
		return
	}
	eventIdStr := c.Param("id")
	c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Event with id '%s' does not exist", eventIdStr)})
}
