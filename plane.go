package fractal

import (
	"math"
)

type Plane struct {
	complexSet ComplexSet
	width      int
	height     int
	iterations int
	values     [][]int
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
	return p.complexSet.Real.Length() / float64(p.width)
}

func (p Plane) YStep() float64 {
	return p.complexSet.Imaginary.Length() / float64(p.height)
}

func (p Plane) Deviation() float64 {
	return deviation(p.values)
}

func (p Plane) Box() Box {
	return Box{0, 0, p.width, p.height}
}

func (p Plane) Value(x, y int) int {
	return p.values[x][y]
}

func deviation(plane [][]int) float64 {
	m := mean(plane)
	sum := 0.0
	count := 0.0
	for _, col := range plane {
		for _, val := range col {
			count++
			sum += math.Pow(float64(val)-m, 2.0)
		}
	}
	return math.Sqrt(sum / count)
}

func mean(plane [][]int) float64 {
	count := 0
	sum := 0
	for _, col := range plane {
		for _, val := range col {
			count++
			sum += val
		}
	}
	return float64(sum) / float64(count)
}

func crop(plane [][]int, box Box) [][]int {
	part := make([][]int, box.Width)
	px := 0
	for x, col := range plane {
		if x >= box.X && x < (box.Width+box.X) {
			partCol := make([]int, box.Height)
			py := 0
			for y, val := range col {
				if y >= box.Y && y < (box.Height+box.Y) {
					partCol[py] = val
					py++
				}
			}
			part[px] = partCol
			px++
		}
	}
	return part
}

func (p Plane) Crop(b Box) Plane {
	xstart := float64(b.X)*p.XStep() + p.complexSet.Real.Start
	xend := float64(b.Width)*p.XStep() + xstart
	ystart := float64(b.Y)*p.YStep() + p.complexSet.Imaginary.Start
	yend := float64(b.Height)*p.YStep() + ystart
	zoomSet := ComplexSet{
		Real:      Range{xstart, xend},
		Imaginary: Range{ystart, yend},
		Algorithm: p.complexSet.Algorithm,
	}
	return Plane{
		complexSet: zoomSet,
		width:      b.Width,
		height:     b.Height,
		iterations: p.iterations,
		values:     crop(p.values, b),
	}
}

func (p Plane) Scale(width int, height int) Plane {
	return p.complexSet.Plane(width, height, p.iterations)
}
