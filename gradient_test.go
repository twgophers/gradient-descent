package gradient

import (
	"fmt"
	"testing"
)

func square(x float64) float64     { return x * x }
func increment(x float64) float64  { return x + 1 }
func derivative(x float64) float64 { return 2 * x }

func TestDifferenceQuotient(t *testing.T) {
	cases := []struct {
		fn        func(float64) float64
		variable  float64
		variation float64
		want      float64
	}{
		{square, 2, 0.0001, 4.0001000000078335},
		{increment, 1, 0.0001, 0.9999999999976694},
		{derivative, 1, 0.0001, 1.9999999999997797},
	}
	for _, c := range cases {
		got := DifferenceQuotient(c.fn, c.variable, c.variation)
		if got != c.want {
			t.Errorf("differenceQuotient(%v, %v, %v) expected: %v but got: %v",
				"square", c.variable, c.variation, c.want, got)
		}
	}
}

func ExampleDifferenceQuotient() {
	fmt.Println(DifferenceQuotient(square, 2, 0.1))
	fmt.Println(DifferenceQuotient(increment, 2, 0.1))
	fmt.Println(DifferenceQuotient(derivative, 2, 0.1))
	//Output:
	//4.100000000000001
	//1.0000000000000009
	//2.0000000000000018
}
