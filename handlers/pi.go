package handlers

import (
	"fmt"
	"math"
	"math-services/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDigitsOfPi(c *gin.Context) {
	// Obtiene el parámetro "digits" del querystring
	digitsStr := c.DefaultQuery("digits", "0")

	// Convierte el string de dígitos a int
	digits, err := strconv.Atoi(digitsStr)
	if err != nil || digits < 0 {
			c.JSON(400, gin.H{
					"error": "El valor de 'digits' debe ser un número no negativo",
			})
			return
	}

	// Limita la cantidad de dígitos a, digamos, 15 para evitar respuestas demasiado largas
	// Especialmente porque la serie de Leibniz no es el método más eficiente para calcular π
	if digits > 15 {
			digits = 15
	}

	// Calcula π usando la serie de Leibniz
	calculatedPi := utils.CalculatePi(digits)
	formattedCalculatedPi := fmt.Sprintf("%.*f", digits, calculatedPi)

	// Obtiene el valor real de π del paquete math
	piStr := fmt.Sprintf("%.*f", digits, math.Pi)

	isMatching := formattedCalculatedPi == piStr

	c.JSON(200, gin.H{
			"digits": digits,
			"calculatedPi": formattedCalculatedPi,
			"realPi": piStr,
			"isMatching": isMatching,
	})
}
