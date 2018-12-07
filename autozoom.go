package fractal

import (
	"math"
)

func (p Plane) AutoZoom() Plane {
	wHalf := p.resolution.Width / 2
	hHalf := p.resolution.Height / 2
	frames := [][]int{
		[]int{0, 0, wHalf, hHalf},
		[]int{wHalf, 0, wHalf, hHalf},
		[]int{wHalf, hHalf, wHalf, hHalf},
		[]int{wHalf, 0, wHalf, hHalf},
	}
	currentDeviation := 0.0
	bestFrame := frames[0]
	for _, f := range frames {
		part := part(p.values, f[0], f[1], f[2], f[3])
		d := deviation(part)
		if d > currentDeviation {
			currentDeviation = d
			bestFrame = f
		}
	}
	xstart := float64(bestFrame[0])*p.XStep() + p.complexSet.Real.Start
	xend := float64(bestFrame[2])*p.XStep() + p.complexSet.Real.Start
	ystart := float64(bestFrame[1])*p.YStep() + p.complexSet.Imaginary.Start
	yend := float64(bestFrame[3])*p.YStep() + p.complexSet.Imaginary.Start
	return p.Zoom(
		Range{xstart, xend},
		Range{ystart, yend},
	)
}

func (p Plane) Zoom(real Range, imaginary Range) Plane {
	zoomSet := ComplexSet{
		Real:      real,
		Imaginary: imaginary,
		Algorithm: p.complexSet.Algorithm,
	}
	return NewPlane(zoomSet, p.resolution, p.iterations)
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

func part(plane [][]int, xoffset int, yoffset int, width int, height int) [][]int {
	part := make([][]int, width)
	for x, col := range plane {
		px := x - xoffset
		if px >= 0 && px < width {
			partCol := make([]int, height)
			for y, val := range col {
				py := y - yoffset
				if py >= 0 && py < height {
					partCol[py] = val
				}
			}
			part[px] = partCol
		}
	}
	return part
}
