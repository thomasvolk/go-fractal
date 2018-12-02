package mandelbrot

import (
	"math/cmplx"
	"sync"
)

type Mandelbrot struct {
	Xstart     float64
	Xend       float64
	Ystart     float64
	Yend       float64
	Iterations int
	Width      int
	Height     int
}

type MandelbrotSet struct {
	Set        [][]int
	Width      int
	Height     int
	Iterations int
}

func (m *Mandelbrot) Set() MandelbrotSet {
	resultSet := make([][]int, m.Height)
	xStep := (m.Xend - m.Xstart) / float64(m.Width)
	yStep := (m.Yend - m.Ystart) / float64(m.Height)

	var wg sync.WaitGroup
	wg.Add(m.Width * m.Height)

	for y := 0; y < m.Height; y++ {
		resultSetRow := make([]int, m.Width)
		for x := 0; x < m.Width; x++ {
			go func(row []int, px int, py int) {
				defer wg.Done()
				i := m.calculateIterations(px, py, xStep, yStep)
				row[px] = i
			}(resultSetRow, x, y)
		}
		resultSet[y] = resultSetRow
	}

	wg.Wait()
	return MandelbrotSet{resultSet, m.Width, m.Height, m.Iterations}
}

func (m *Mandelbrot) calculateIterations(x int, y int, xStep float64, yStep float64) int {
	c := complex((float64(x)*xStep)+m.Xstart,
		(float64(y)*yStep)+m.Ystart)
	z := 0 + 0i
	i := 0
	for ; i < m.Iterations; i++ {
		z = z*z + c
		if cmplx.Abs(z) >= 2 {
			break
		}
	}
	return i
}
