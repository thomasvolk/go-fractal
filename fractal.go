package fractal

import "sync"

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

func (r Range) Add(o Range) Range {
	return Range{r.Start + o.Start, r.End + o.End}
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

func (p Plane) Crop(x int, y int, width int, height int) Plane {
	xstart := float64(x)*p.XStep() + p.complexSet.Real.Start
	xend := float64(width)*p.XStep() + xstart
	ystart := float64(y)*p.YStep() + p.complexSet.Imaginary.Start
	yend := float64(height)*p.YStep() + ystart
	zoomSet := ComplexSet{
		Real:      Range{xstart, xend},
		Imaginary: Range{ystart, yend},
		Algorithm: p.complexSet.Algorithm,
	}
	return NewPlane(zoomSet, p.width, p.height, p.iterations)
}
