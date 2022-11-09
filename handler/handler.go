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

func (h *Handler) ArithmeticOperation(c *gin.Context) {

	var newRequest models.OperationRequest
	if err := c.ShouldBind(&newRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var result int64
	var operation models.OperationType

	switch {
	case newRequest.OperationType == "addition":
		result = newRequest.X + newRequest.Y
		operation = models.Addition

	case newRequest.OperationType == "subtraction":
		result = newRequest.X - newRequest.Y
		operation = models.Subtraction

	case newRequest.OperationType == "multiplication":
		result = newRequest.X * newRequest.Y
		operation = models.Multiplication
	}

	myResponse := models.OperationResponse{
		"BemEmma",
		result,
		operation.String(),
	}
	c.IndentedJSON(http.StatusOK, myResponse)
}
