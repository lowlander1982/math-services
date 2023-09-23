package models

type MathFormula struct {
	MathFunction string             `json:"math_function"`
	Variables    map[string]float64 `json:"variables"`
	Constants    []string           `json:"constants"`
	Complex      bool               `json:"complex"`
}
