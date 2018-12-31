package fractal

import (
	"image"
	"math"
)

type Value struct {
	X     int
	Y     int
	Value int
}

func (p Plane) Shape(x int, y int, angleStep float64) []image.Point {
	var points []image.Point
	for angle := 0.0; angle < 360.0; angle += angleStep {
		axis := p.axis(x, y, angle)
		pitch := maxPitch(axis)
		points = append(points, image.Point{pitch.X, pitch.Y})
	}
	return points
}

func maxPitch(axis []Value) Value {
	result := axis[0]
	lastPitch := 0
	last := axis[0]
	for _, current := range axis {
		p := abs(current.Value - last.Value)
		if p > lastPitch {
			lastPitch = p
			result = current
		}
		last = current
	}
	return result
}

func abs(v int) int {
	if v < 0 {
		return v * -1
	}
	return v
}

func (p Plane) axis(x int, y int, angle float64) []Value {
	box := p.Box().InnerBox(x, y)
	var axis []Value
	px := x
	py := y
	for r := float64(1); px < box.Width && py < box.Height; r++ {
		value := Value{px, py, p.Value(px, py)}
		axis = append(axis, value)
		px = x + int(r*math.Cos(angle))
		py = y + int(r*math.Sin(angle))
	}
	return axis
}
