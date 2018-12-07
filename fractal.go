package fractal

type ComplexSet struct {
	Real      Range
	Imaginary Range
	Algorithm func(x float64, y float64, iterations int) int
}

type Range struct {
	Start float64
	End   float64
}

type Resolution struct {
	Width  int
	Height int
}

type Plane struct {
	complexSet ComplexSet
	resolution Resolution
	iterations int
	values     [][]int
}

func (r Range) Add(o Range) Range {
	return Range{r.Start + o.Start, r.End + o.End}
}

func (p Plane) XStep() float64 {
	return (p.complexSet.Real.End - p.complexSet.Real.Start) / float64(p.resolution.Width)
}

func (p Plane) YStep() float64 {
	return (p.complexSet.Imaginary.End - p.complexSet.Imaginary.Start) / float64(p.resolution.Height)
}

func NewPlane(complexSet ComplexSet, resolution Resolution, iterations int) Plane {
	p := Plane{
		complexSet: complexSet,
		resolution: resolution,
		iterations: iterations,
		values:     make([][]int, resolution.Width),
	}
	xStep := p.XStep()
	yStep := p.YStep()

	for x := 0; x < p.resolution.Width; x++ {
		col := make([]int, p.resolution.Height)
		for y := 0; y < p.resolution.Height; y++ {
			cx := (float64(x) * xStep) + p.complexSet.Real.Start
			cy := (float64(y) * yStep) + p.complexSet.Imaginary.Start
			count := p.complexSet.Algorithm(cx, cy, p.iterations)
			col[y] = count
		}
		p.values[x] = col
	}
	return p
}
