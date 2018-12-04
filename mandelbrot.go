package fractal

import (
	"math/cmplx"
	"sync"
)

func Mandelbrot(cs ComplexSet) IterationMapping {
	r := cs.Resolution
	xrange := cs.XRange
	yrange := cs.YRange
	resultSet := make([][]int, r.Height)
	xStep := (xrange.End - xrange.Start) / float64(r.Width)
	yStep := (yrange.End - xrange.Start) / float64(r.Height)

	var wg sync.WaitGroup
	wg.Add(r.Width * r.Height)

	for y := 0; y < r.Height; y++ {
		resultSetRow := make([]int, r.Width)
		for x := 0; x < r.Width; x++ {
			go func(fcs ComplexSet, row []int, px int, py int) {
				defer wg.Done()
				i := calculateIterations(fcs, px, py, xStep, yStep)
				row[px] = i
			}(cs, resultSetRow, x, y)
		}
		resultSet[y] = resultSetRow
	}

	wg.Wait()
	return NewIterationMapping(resultSet)
}

func calculateIterations(cs ComplexSet, x int, y int, xStep float64, yStep float64) int {
	c := complex((float64(x)*xStep)+cs.XRange.Start,
		(float64(y)*yStep)+cs.YRange.Start)
	z := 0 + 0i
	i := 0
	for ; i < cs.Iterations; i++ {
		z = z*z + c
		if cmplx.Abs(z) >= 2 {
			break
		}
	}
	return i
}
