package fractal

import (
	"math"
)

func (p Plane) CircleAutoZoom(x int, y int, radiusDivisor float64, angleStep float64) Plane {
	box := innerBox(x, y, Box{0, 0, p.width, p.height})
	rx := float64(box.width) / radiusDivisor
	ry := float64(box.height) / radiusDivisor
	currentDeviation := 0.0
	bestFrame := box
	for theta := 0.0; theta < 360.0; theta += angleStep {
		px := x + int(rx*math.Cos(theta))
		py := y + int(ry*math.Sin(theta))
		frame := innerBox(px, py, box)
		part := part(p.values, frame)
		d := deviation(part)
		if d > currentDeviation {
			currentDeviation = d
			bestFrame = frame
		}
	}
	return p.Crop(bestFrame)
}

func (p Plane) RasterAutoZoom(division int) Plane {
	wd := p.width / division
	hd := p.height / division

	boxes := division * division
	frames := make([]Box, boxes)
	for f := 0; f < boxes; f++ {
		x := (f * wd) % p.width
		y := (f * hd) % p.height
		frames[f] = Box{x, y, wd, hd}
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
