package app

import (
	"deploy-sample/src/controllers"
	"deploy-sample/src/logger"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func mapUrls() {
	router.GET("/ping", func(c *gin.Context) {
		logger.Info("logger is working which is using a package")
		new := os.Getenv("TEST_ENV")
		fmt.Println(new)
		c.JSON(http.StatusOK, map[string]string{
			"status": "working",
			"ENV":    new,
		})
		return
	})
	router.GET("/cont", controllers.TestController)
}
