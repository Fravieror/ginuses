package main

import (
	"fmt"

	"gintest/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("hello")
	router := gin.Default()

	handlers := handler.Handler{}
	router.GET("/hot-100/:date", handlers.GetHot100)
	router.POST("/save-album", handlers.PostAlbums)
	_ = router.Run(fmt.Sprintf("localhost:%s", "8080"))
}
