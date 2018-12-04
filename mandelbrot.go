package fractal

import (
	"math/cmplx"
	"sync"
)

func Mandelbrot(conf Config) IterationMapping {
	r := conf.Resolution
	xrange := conf.Real
	yrange := conf.Imaginary
	resultSet := make([][]int, r.Height)
	xStep := (xrange.End - xrange.Start) / float64(r.Width)
	yStep := (yrange.End - yrange.Start) / float64(r.Height)

	var wg sync.WaitGroup
	wg.Add(r.Width * r.Height)

	for y := 0; y < r.Height; y++ {
		resultSetRow := make([]int, r.Width)
		for x := 0; x < r.Width; x++ {
			go func(fcs Config, row []int, px int, py int) {
				defer wg.Done()
				i := calculateIterations(fcs, px, py, xStep, yStep)
				row[px] = i
			}(conf, resultSetRow, x, y)
		}
		resultSet[y] = resultSetRow
	}

	wg.Wait()
	return NewIterationMapping(resultSet)
}

func calculateIterations(conf Config, x int, y int, xStep float64, yStep float64) int {
	c := complex((float64(x)*xStep)+conf.Real.Start,
		(float64(y)*yStep)+conf.Imaginary.Start)
	z := 0 + 0i
	i := 0
	for ; i < conf.Iterations; i++ {
		z = z*z + c
		if cmplx.Abs(z) >= 2 {
			break
		}
	}
	return i
}
