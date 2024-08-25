package main

import (
	"gin/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/athlete/:id", controllers.GetAthlete)

	router.Run()
}
