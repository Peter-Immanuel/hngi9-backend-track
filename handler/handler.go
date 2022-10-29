package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"

	models "github.com/Peter-Immanuel/hngi9-backend-track/models"
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
	myProfile := models.NewBioProfile()
	myProfile.CreateMyBio()
	c.IndentedJSON(http.StatusOK, myProfile)
	return
}
