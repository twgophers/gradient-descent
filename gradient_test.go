package gradient

import (
	"fmt"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
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

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Approximating a derivative with a difference quotient"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p,
		"Function", addDerivative(-10, 10, square),
		"Actual", addDerivative(-10, 10, derivative),
		"Estimate", addDerivativeEstimative(-10, 10),
	)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "derivatives.png"); err != nil {
		panic(err)
	}
}

func addDerivative(start, end int, fn func(float64) float64) plotter.XYs {
	pts := make(plotter.XYs, end-start)
	for i := range pts {
		pts[i].X = float64(i)
		pts[i].Y = float64(fn(pts[i].X))
	}
	return pts
}

func addDerivativeEstimative(start, end float64) plotter.XYs {
	pts := make(plotter.XYs, int(end-start))
	for i := range pts {
		pts[i].X = float64(i)
		pts[i].Y = DifferenceQuotient(square, pts[i].X, 0.1)
	}
	return pts
}
