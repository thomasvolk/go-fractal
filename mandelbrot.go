package fractal

import (
	"math/cmplx"
)

func Mandelbrot(x float64, y float64, iterations int) int {
	c := complex(x, y)
	z := 0 + 0i
	i := 0
	for ; i < iterations; i++ {
		z = z*z + c
		if cmplx.Abs(z) >= 2 {
			break
		}
	}
	return i
}
