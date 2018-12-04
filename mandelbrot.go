package fractal

import (
	"math/cmplx"
)

func Mandelbrot(conf Config) [][]int {
	r := conf.Resolution
	xStep := conf.XStep()
	yStep := conf.YStep()

	result := make([][]int, r.Width)

	for x := 0; x < r.Width; x++ {
		col := make([]int, r.Height)
		for y := 0; y < r.Height; y++ {
			cx := (float64(x) * xStep) + conf.Real.Start
			cy := (float64(y) * yStep) + conf.Imaginary.Start
			count := doRecursion(cx, cy, conf.Iterations)
			col[y] = count
		}
		result[x] = col
	}
	return result
}

func doRecursion(x float64, y float64, iterations int) int {
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
