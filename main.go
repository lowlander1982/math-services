package main

import (
	"math-services/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/pi", handlers.GetDigitsOfPi)

	r.POST("/math_function", handlers.ParseMathFormula)
	r.Run(":8080")
}
