package fractal

import (
	"math"

	model "./model"
	utils "./utils"
)

func (p Plane) CircleAutoZoom(x int, y int, radiusDivisor float64, angleStep float64) Plane {
	return p.CircleZoom(utils.Deviation, x, y, radiusDivisor, angleStep)
}

func (p Plane) CircleZoom(selector func(part [][]int) float64, x int, y int, radiusDivisor float64, angleStep float64) Plane {
	box := model.Box{0, 0, p.width, p.height}.InnerBox(x, y)
	rx := float64(box.Width) / radiusDivisor
	ry := float64(box.Height) / radiusDivisor
	currentDeviation := 0.0
	bestFrame := box
	for theta := 0.0; theta < 360.0; theta += angleStep {
		px := x + int(rx*math.Cos(theta))
		py := y + int(ry*math.Sin(theta))
		frame := box.InnerBox(px, py)
		part := utils.Crop(p.values, frame)
		d := selector(part)
		if d > currentDeviation {
			currentDeviation = d
			bestFrame = frame
		}
	}
	return p.Crop(bestFrame)
}
func (p Plane) RasterAutoZoom(division int) Plane {
	return p.RasterZoom(utils.Deviation, division)
}

func (p Plane) RasterZoom(selector func(part [][]int) float64, division int) Plane {
	wd := p.width / division
	hd := p.height / division

	boxes := division * division
	frames := make([]model.Box, boxes)
	for f := 0; f < boxes; f++ {
		x := (f * wd) % p.width
		y := (f * hd) % p.height
		frames[f] = model.Box{x, y, wd, hd}
	}

	currentDeviation := 0.0
	bestFrame := frames[0]
	for _, f := range frames {
		part := utils.Crop(p.values, f)
		d := utils.Deviation(part)
		if d > currentDeviation {
			currentDeviation = d
			bestFrame = f
		}
	}
	return p.Crop(bestFrame)
}
