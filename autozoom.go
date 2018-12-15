package fractal

import (
	"math"
)

func (p Plane) AutoZoom() Plane {
	wHalf := p.width / 2
	hHalf := p.height / 2
	frames := []Box{
		Box{0, 0, wHalf, hHalf},
		Box{wHalf, 0, wHalf, hHalf},
		Box{wHalf, hHalf, wHalf, hHalf},
		Box{wHalf, 0, wHalf, hHalf},
	}
	currentDeviation := 0.0
	bestFrame := frames[0]
	for _, f := range frames {
		part := part(p.values, f)
		d := deviation(part)
		if d > currentDeviation {
			currentDeviation = d
			bestFrame = f
		}
	}
	return p.Crop(bestFrame)
}

func innerBox(x int, y int, outer Box) Box {
	nx := (x - outer.x)
	ny := (y - outer.y)
	xmin := (outer.width - nx)
	ymin := (outer.height - ny)
	if nx < xmin {
		xmin = nx
	}
	if ny < ymin {
		ymin = ny
	}

	widthScaleFactor := float64(xmin*2) / float64(outer.height)
	heightScaleFactor := float64(ymin*2) / float64(outer.height)

	scaleFactor := heightScaleFactor
	if widthScaleFactor < heightScaleFactor {
		scaleFactor = widthScaleFactor
	}

	newWidth := int(float64(outer.width) * scaleFactor)
	newHeight := int(float64(outer.height) * scaleFactor)

	return Box{x - newWidth/2, y - newHeight/2, newWidth, newHeight}
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

func part(plane [][]int, box Box) [][]int {
	part := make([][]int, box.width)
	for x, col := range plane {
		px := x - box.x
		if px >= 0 && px < box.width {
			partCol := make([]int, box.height)
			for y, val := range col {
				py := y - box.y
				if py >= 0 && py < box.height {
					partCol[py] = val
				}
			}
			part[px] = partCol
		}
	}
	return part
}
