package handler

import (
	"github.com/gin-gonic/gin"
)

type MessageResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, MessageResponse{message})
}

func newSuccessResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, MessageResponse{message})
}
