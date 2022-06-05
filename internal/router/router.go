package router

import (
	"tusa/internal/events"
	"tusa/internal/storage"

	"github.com/gin-gonic/gin"
)

type WebServer interface {
	Run(addr ...string) (err error)
}

func SetupServer() WebServer {
	router := gin.Default()
	tusaStorage := storage.NewTusaEventStore()
	rsvpStorage := storage.NewRsvpStore()
	eventHandler := events.NewEventHandler(tusaStorage, rsvpStorage)
	router.GET("/events", eventHandler.List)
	router.POST("/events", eventHandler.Add)
	router.GET("/events/:id", eventHandler.Get)
	router.POST("/events/:id/rsvp", eventHandler.Rsvp)
	return router
}
