package fractal

import (
	"fmt"
	"sync"
)

type Box struct {
	x      int
	y      int
	width  int
	height int
}

func (b Box) String() string {
	return fmt.Sprintf("Box(x=%d, y=%d, w=%d, h=%d)", b.x, b.y, b.width, b.height)
}

type ComplexSet struct {
	Real      Range
	Imaginary Range
	Algorithm func(x float64, y float64, iterations int) int
}

type Range struct {
	Start float64
	End   float64
}

type Plane struct {
	complexSet ComplexSet
	width      int
	height     int
	iterations int
	values     [][]int
}

func (r Range) String() string {
	return fmt.Sprintf("%v-%v", r.Start, r.End)
}

func (p Plane) Width() int {
	return p.width
}

func (p Plane) Height() int {
	return p.height
}

func (p Plane) ComplexSet() ComplexSet {
	return p.complexSet
}

func (p Plane) XStep() float64 {
	return (p.complexSet.Real.End - p.complexSet.Real.Start) / float64(p.width)
}

func (p Plane) YStep() float64 {
	return (p.complexSet.Imaginary.End - p.complexSet.Imaginary.Start) / float64(p.height)
}

func NewPlane(complexSet ComplexSet, width int, heigth int, iterations int) Plane {
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
			go p.recursion(&wg, col, y, cx, cy)

		}
		p.values[x] = col
	}
	wg.Wait()
	return p
}

func (p Plane) recursion(wg *sync.WaitGroup, col []int, y int, cx float64, cy float64) {
	defer wg.Done()
	count := p.complexSet.Algorithm(cx, cy, p.iterations)
	col[y] = count
}

func (p Plane) Crop(b Box) Plane {
	xstart := float64(b.x)*p.XStep() + p.complexSet.Real.Start
	xend := float64(b.width)*p.XStep() + xstart
	ystart := float64(b.y)*p.YStep() + p.complexSet.Imaginary.Start
	yend := float64(b.height)*p.YStep() + ystart
	zoomSet := ComplexSet{
		Real:      Range{xstart, xend},
		Imaginary: Range{ystart, yend},
		Algorithm: p.complexSet.Algorithm,
	}
	return NewPlane(zoomSet, p.width, p.height, p.iterations)
}
