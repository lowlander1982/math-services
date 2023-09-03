package main

import (
	"math-services/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/pi", handlers.GetDigitsOfPi)
	r.Run(":8080")
}
