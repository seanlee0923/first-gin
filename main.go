package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Gin 엔진 생성
	router := gin.Default()

	// 라우팅 정의
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, this is Gin!",
		})
	})

	// 웹 서버 시작
	router.Run(":8080")
}
