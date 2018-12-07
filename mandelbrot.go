package fractal

import (
	"math/cmplx"
)

type Mandelbrot struct {
	real      Range
	imaginary Range
}

func NewMandelbrot(real Range, imaginary Range) Mandelbrot {
	return Mandelbrot{real, imaginary}
}

func (m Mandelbrot) Real() Range { return m.real }

func (m Mandelbrot) Imaginary() Range { return m.imaginary }

func (m Mandelbrot) doRecursion(x float64, y float64, iterations int) int {
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
