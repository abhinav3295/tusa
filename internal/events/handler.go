package events

import "github.com/gin-gonic/gin"

type TusaEventHandler interface {
	List(c *gin.Context)
	Add(c *gin.Context)
}

type tusaEventHandler struct {
	events []TusaEvent
}

func NewEventHandler() TusaEventHandler {
	return &tusaEventHandler{
		events: make([]TusaEvent, 20),
	}
}

func (h tusaEventHandler) List(c *gin.Context) {

}
func (h *tusaEventHandler) Add(c *gin.Context) {

}
