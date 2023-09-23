package constants

import "math"

var Constants = map[string]float64{
	"pi": math.Pi,
	"g":  9.81,
	"k":  8.8541878128 * math.Pow(10, -12),
}
