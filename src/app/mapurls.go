package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func mapUrls() {
	router.GET("/ping", func(c *gin.Context) {
		new := os.Getenv("TEST_ENV")
		fmt.Println(new)
		c.JSON(http.StatusOK, map[string]string{
			"status": "working",
			"ENV":    new,
		})
		return
	})

}
