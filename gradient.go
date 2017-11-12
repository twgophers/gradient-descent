package gradient

import (
	"github.com/twgophers/collections"
)

//
func DifferenceQuotient(fn func(float64) float64, x, h float64) float64 {
	return (fn(x+h) - fn(x)) / h
}

func SumOfSquares(vector collections.Vector) collections.Vector {
	sum := collections.Vector{}
	for _, c := range vector {
		sum = append(sum, 2*c)
	}
	return sum
}
