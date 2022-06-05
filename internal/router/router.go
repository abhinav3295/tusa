package router

import (
	"tusa/internal/events"

	"github.com/gin-gonic/gin"
)

type WebServer interface {
	Run(addr ...string) (err error)
}

func SetupServer() WebServer {
	router := gin.Default()
	eventHandler := events.NewEventHandler()
	router.GET("/events", eventHandler.List)
	router.POST("/events", eventHandler.Add)
	return router
}
