package simple_expresions

import (
	"fmt"
	"math-services/controllers/constants"
	"math-services/models"

	"github.com/Knetic/govaluate"
)

func MathParser(formula models.MathFormula, complexFunctions map[string]govaluate.ExpressionFunction) (interface{}, error) {
	var expr *govaluate.EvaluableExpression
	var err error
	if complexFunctions != nil {
		expr, err = govaluate.NewEvaluableExpressionWithFunctions(formula.MathFunction, complexFunctions)
	} else {
		expr, err = govaluate.NewEvaluableExpression(formula.MathFunction)
	}

	if err != nil {
		return 0, err
	}

	var params map[string]interface{} = make(map[string]interface{})

	for k, v := range formula.Variables {
		params[k] = v
	}

	for _, v := range formula.Constants {
		constant, ok := constants.Constants[v]
		if !ok {
			return 0, fmt.Errorf("constant %s doesn't exist", v)
		}
		params[v] = constant
	}

	result, err := expr.Evaluate(params)

	if err != nil {
		return 0, err
	}

	return result, nil
}
