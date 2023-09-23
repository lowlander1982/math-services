package main

import (
	"math"
	"math-services/controllers/complex_expresions"
	"math-services/controllers/simple_expresions"
	"math-services/handlers"
	"math-services/models"
	"math-services/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/pi", handlers.GetDigitsOfPi)
	return r
}

func TestDefaultDigits(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/pi", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "\"realPi\":\"3\"")
}

func TestTwoDigits(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/pi?digits=2", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "\"realPi\":\"3.14\"")
}

func TestInvalidDigits(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/pi?digits=abc", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
	assert.Contains(t, resp.Body.String(), "El valor de 'digits' debe ser un número no negativo")
}

func TestNegativeDigits(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/pi?digits=-5", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
	assert.Contains(t, resp.Body.String(), "El valor de 'digits' debe ser un número no negativo")
}

func TestLargeDigits(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/pi?digits=1000", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "\"realPi\":\"3.141592653589793\"")
}

func TestCalculatePiFunctionPass(t *testing.T) {
	// Se esta calculando con un metodo muy poco preciso, por lo que se usa un delta de 0.01
	assert.InDelta(t, math.Pi, utils.CalculatePi(15), 0.01)
}

func TestMathParserComplexFails(t *testing.T) {
	formula := models.MathFormula{
		MathFunction: "sin(pi)",
		Variables:    map[string]float64{},
		Constants:    []string{"pi"},
		Complex:      false,
	}
	result, err := simple_expresions.MathParser(formula, nil)
	assert.Equal(t, result, 0)
	assert.Error(t, err)
}
func TestMathParserComplex(t *testing.T) {
	formula := models.MathFormula{
		MathFunction: "sin(pi)",
		Variables:    map[string]float64{},
		Constants:    []string{"pi"},
		Complex:      true,
	}
	result, err := simple_expresions.MathParser(formula, complex_expresions.GetComplexFunctions())
	assert.Equal(t, math.Sin(math.Pi), result)
	assert.Equal(t, nil, err)
}
func TestMathParserConstantsFails(t *testing.T) {
	formula := models.MathFormula{
		MathFunction: "3*pi",
		Variables:    map[string]float64{},
		Constants:    []string{},
		Complex:      false,
	}
	result, err := simple_expresions.MathParser(formula, nil)
	assert.Equal(t, 0, result)
	assert.Error(t, err)
}
func TestMathParserConstants(t *testing.T) {
	formula := models.MathFormula{
		MathFunction: "3*pi",
		Variables:    map[string]float64{},
		Constants:    []string{"pi"},
		Complex:      false,
	}
	result, err := simple_expresions.MathParser(formula, nil)
	assert.Equal(t, 3*math.Pi, result)
	assert.Equal(t, nil, err)
}
func TestMathParserVariablesFails(t *testing.T) {
	formula := models.MathFormula{
		MathFunction: "a*b",
		Variables: map[string]float64{
			"a": 4,
		},
		Constants: []string{},
		Complex:   false,
	}
	result, err := simple_expresions.MathParser(formula, nil)
	assert.Equal(t, 0, result)
	assert.Error(t, err)
}
func TestMathParserVariables(t *testing.T) {
	formula := models.MathFormula{
		MathFunction: "a*b",
		Variables: map[string]float64{
			"a": 4,
			"b": 3,
		},
		Constants: []string{},
		Complex:   false,
	}
	result, err := simple_expresions.MathParser(formula, nil)
	assert.Equal(t, 12.0, result)
	assert.Equal(t, nil, err)
}
func TestMathParserSimple(t *testing.T) {
	formula := models.MathFormula{
		MathFunction: "4*3",
		Variables:    map[string]float64{},
		Constants:    []string{},
		Complex:      false,
	}
	result, err := simple_expresions.MathParser(formula, nil)
	assert.Equal(t, 12.0, result)
	assert.Equal(t, nil, err)
}
func TestMathParserTan(t *testing.T) {
	formula := models.MathFormula{
		MathFunction: "tan(pi/3)",
		Variables:    map[string]float64{},
		Constants:    []string{"pi"},
		Complex:      true,
	}
	result, err := simple_expresions.MathParser(formula, nil)
	assert.Equal(t, math.Tan(math.Pi/3), result)
	assert.Equal(t, nil, err)
}
