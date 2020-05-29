package controllers

import (
	"deploy-sample/src/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TestController : this is a test controller
func TestController(c *gin.Context) {
	logger.Info("logger is working in controller")

	c.JSON(http.StatusOK, map[string]string{
		"status": "controller is working",
	})
	return
}
