package fractal

import "math"

func (p Plane) CircleAutoZoom(x int, y int, radiusDivisor float64, angleStep float64) Plane {
	frames := p.CircleFrames(x, y, radiusDivisor, angleStep)
	currentDeviation := 0.0
	bestFrame := p
	for _, frame := range frames {
		d := frame.Deviation()
		if d > currentDeviation {
			currentDeviation = d
			bestFrame = frame
		}

	}
	return bestFrame
}

func (p Plane) CircleFrames(x int, y int, radiusDivisor float64, angleStep float64) []Plane {
	box := p.Box().InnerBox(x, y)
	rx := float64(box.Width) / radiusDivisor
	ry := float64(box.Height) / radiusDivisor
	var frames []Plane
	for theta := 0.0; theta < 360.0; theta += angleStep {
		px := x + int(rx*math.Cos(theta))
		py := y + int(ry*math.Sin(theta))
		frames = append(frames, p.Crop(box.InnerBox(px, py)))
	}
	return frames
}

func (p Plane) RasterAutoZoom(division int) Plane {
	frames := p.RasterFrames(division)
	currentDeviation := 0.0
	bestFrame := p
	for _, f := range frames {
		d := f.Deviation()
		if d > currentDeviation {
			currentDeviation = d
			bestFrame = f
		}
	}
	return bestFrame
}

func (p Plane) RasterFrames(division int) []Plane {
	wd := p.width / division
	hd := p.height / division
	boxes := division * division
	frames := make([]Plane, boxes)
	for f := 0; f < boxes; f++ {
		x := (f * wd) % p.width
		y := (f * hd) % p.height
		frames[f] = p.Crop(Box{x, y, wd, hd})
	}
	return frames
}
