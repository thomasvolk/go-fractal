package fractal

import (
	"sync"
)

type ComplexSet struct {
	Real      Range
	Imaginary Range
	Algorithm func(x float64, y float64, iterations int) int
}

func (complexSet ComplexSet) Plane(width int, heigth int, iterations int) Plane {
	p := Plane{
		complexSet: complexSet,
		width:      width,
		height:     heigth,
		iterations: iterations,
		values:     make([][]int, width),
	}
	xStep := p.XStep()
	yStep := p.YStep()

	var wg sync.WaitGroup
	wg.Add(p.width * p.height)

	for x := 0; x < p.width; x++ {
		col := make([]int, p.height)
		for y := 0; y < p.height; y++ {
			cx := (float64(x) * xStep) + p.complexSet.Real.Start
			cy := (float64(y) * yStep) + p.complexSet.Imaginary.Start
			go complexSet.doIterations(&wg, col, y, cx, cy, iterations)

		}
		p.values[x] = col
	}
	wg.Wait()
	return p
}

func (complexSet ComplexSet) doIterations(wg *sync.WaitGroup, col []int, y int, cx float64, cy float64, iterations int) {
	defer wg.Done()
	count := complexSet.Algorithm(cx, cy, iterations)
	col[y] = count
}
