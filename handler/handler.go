package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	ctx context.Context
}

func NewHandler(context context.Context) *Handler {
	return &Handler{
		context,
	}
}

func (h *Handler) BioHandler(c *gin.Context) {
	c.IndentedJson(http.StatusOK)
	return
}
