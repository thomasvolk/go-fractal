package fractal

import "fmt"

type Range struct {
	Start float64
	End   float64
}

func (r Range) String() string {
	return fmt.Sprintf("C%v-R%v", r.Center(), r.Radius())
}

func (r Range) Center() float64 {
	return r.Start + r.Radius()
}

func (r Range) Radius() float64 {
	return r.Length() / 2.0
}

func (r Range) Length() float64 {
	return r.End - r.Start
}

func NewRange(center float64, radius float64) Range {
	return Range{center - radius, center + radius}
}
