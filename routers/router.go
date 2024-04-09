package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/posts", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hi",
		})
	})

	return router
}
