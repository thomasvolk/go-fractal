package fractal

import (
	"math"
)

func (p Plane) Deviation() float64 {
	return deviation(p.values)
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
	for x, col := range plane {
		px := x - box.X
		if px >= 0 && px < box.Width {
			partCol := make([]int, box.Height)
			for y, val := range col {
				py := y - box.Y
				if py >= 0 && py < box.Height {
					partCol[py] = val
				}
			}
			part[px] = partCol
		}
	}
	return part
}
