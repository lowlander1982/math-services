package handlers

import (
	"math-services/controllers/complex_expresions"
	"math-services/controllers/simple_expresions"
	"math-services/models"

	"github.com/Knetic/govaluate"
	"github.com/gin-gonic/gin"
)

func ParseMathFormula(c *gin.Context) {
	var body models.MathFormula
	var complexExpressions map[string]govaluate.ExpressionFunction
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "unable to parse body",
			"error":   err.Error(),
		})
		return
	}
	if body.Complex {
		complexExpressions = complex_expresions.GetComplexFunctions()
	}
	result, err := simple_expresions.MathParser(body, complexExpressions)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "unable to parse formula",
			"error":   err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"result": result,
	})
}
