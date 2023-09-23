package complex_expresions

import (
	"math"

	"github.com/Knetic/govaluate"
)

func GetComplexFunctions() map[string]govaluate.ExpressionFunction {
	return map[string]govaluate.ExpressionFunction{
		"sin": sin,
	}

}

func sin(args ...interface{}) (interface{}, error) {
	return math.Sin(args[0].(float64)), nil
}
