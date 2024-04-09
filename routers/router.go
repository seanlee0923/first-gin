package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	return router
}
