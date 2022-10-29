package main

import (
	handler "github.com/Peter-Immanuel/hngi9-backend-track/handler"
	"github.com/gin-gonic/gin"
)

var bioHandler *handler.Handler

func main() {
	router := gin.Default()
	router.GET("/bio", bioHandler.BioHandler)

	router.Run()
}
