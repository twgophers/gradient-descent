package gradient

//type function func(v1 float64) float64

func DifferenceQuotient(fn func(float64) float64, x, h float64) float64 {
	return (fn(x+h) - fn(x)) / h
}
