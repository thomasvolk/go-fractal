package fractal

import (
	"math"
)

func (p Plane) AutoZoom() Plane {
	wHalf := p.width / 2
	hHalf := p.height / 2
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

func innerBox(x int, y int, width int, height int) (int, int, int, int) {
	xmin := (width - x)
	ymin := (height - y)
	if x < xmin {
		xmin = x
	}
	if y < ymin {
		ymin = y
	}
	widthScaleFactor := float64(xmin*2) / float64(height)
	heightScaleFactor := float64(ymin*2) / float64(height)
	scaleFactor := 0.0
	if widthScaleFactor < heightScaleFactor {
		scaleFactor = widthScaleFactor
	} else {
		scaleFactor = heightScaleFactor
	}
	newWidth := int(float64(width) * scaleFactor)
	newHeight := int(float64(height) * scaleFactor)

	return x - newWidth/2, y - newHeight/2, newWidth, newHeight
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
