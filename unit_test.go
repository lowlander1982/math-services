package main

import (
	"math"
	"math-services/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePiFunction(t *testing.T) {
	// Se esta calculando con un metodo muy poco preciso, por lo que se usa un delta de 0.01
	assert.InDelta(t, math.Pi, utils.CalculatePi(15), 0.01)
}
