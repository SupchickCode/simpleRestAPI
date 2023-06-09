package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string `json:""`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)

	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}
