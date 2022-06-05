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
	eventHandler := events.NewEventHandler(&tusaStorage)
	router.GET("/events", eventHandler.List)
	router.POST("/events", eventHandler.Add)
	return router
}
