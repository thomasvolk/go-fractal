package fractal

import (
	"math"
)

func (p Plane) AutoZoom() Plane {
	wHalf := p.grid.Width / 2
	hHalf := p.grid.Height / 2
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
	return p.Crop(bestFrame[0], bestFrame[1], bestFrame[2], bestFrame[3])
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
