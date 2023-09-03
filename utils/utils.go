package utils

// calculatePi utiliza la serie de Leibniz para aproximar el valor de π.
// No es el método más eficiente, pero es simple de implementar para propósitos demostrativos.
func CalculatePi(digits int) float64 {
	var pi float64
	var operator float64 = 1

	for i := 0; i < digits*10; i++ {
		pi += operator * (4 / (2*float64(i) + 1))
		operator = -operator
	}

	return pi
}
