package fractal

import (
	"math"
)

func (p Plane) CircleAutoZoom(x int, y int, radiusDivisor float64, angleStep float64) Plane {
	box := Box{0, 0, p.width, p.height}.InnerBox(x, y)
	rx := float64(box.Width) / radiusDivisor
	ry := float64(box.Height) / radiusDivisor
	currentDeviation := 0.0
	bestFrame := p
	for theta := 0.0; theta < 360.0; theta += angleStep {
		px := x + int(rx*math.Cos(theta))
		py := y + int(ry*math.Sin(theta))
		frame := p.Crop(box.InnerBox(px, py))
		d := frame.Deviation()
		if d > currentDeviation {
			currentDeviation = d
			bestFrame = frame
		}
	}
	return bestFrame
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
	bestFrame := p
	for _, f := range frames {
		frame := p.Crop(f)
		d := frame.Deviation()
		if d > currentDeviation {
			currentDeviation = d
			bestFrame = frame
		}
	}
	return bestFrame
}
