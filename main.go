package main

import (
	"github.com/seanlee0923/first-gin/routers"
)

func main() {
	// 기본 라우터 생성
	router := routers.DefaultRouter()

	// 웹 서버 시작
	router.Run(":8080")
}
