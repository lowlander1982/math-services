package complex_expresions

import (
	"math"

	"github.com/Knetic/govaluate"
)

func GetComplexFunctions() map[string]govaluate.ExpressionFunction {
	return map[string]govaluate.ExpressionFunction{
		"sin": sin,
		"cos": cos,
		"tan": tan,
		"sqrt": sqrt,
		"abs": abs,
		"exp": exp,
		"log": log,
	}

}

func sin(args ...interface{}) (interface{}, error) {
	return math.Sin(args[0].(float64)), nil
}

func cos(args ...interface{}) (interface{}, error) {
	return math.Cos(args[0].(float64)), nil
}
func sqrt(args ...interface{}) (interface{}, error) {
	return math.Sqrt(args[0].(float64)), nil
}
func abs(args ...interface{}) (interface{}, error) {
	return math.Abs(args[0].(float64)), nil
}
func exp(args ...interface{}) (interface{}, error) {
	return math.Exp(args[0].(float64)), nil
}
func log(args ...interface{}) (interface{}, error) {
	return math.Log(args[0].(float64)), nil
}
func tan(args ...interface{}) (interface{}, error) {
	return math.Tan(args[0].(float64)), nil
}

