package main

import (
	"github.com/gin-gonic/gin"
)

var handler *Handler = Handler.New()

func main() {
	router = gin.Default()
	router.GET("/bio", handler.BioHandler)
}
