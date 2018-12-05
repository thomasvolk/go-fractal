package fractal

type Range struct {
	Start float64
	End   float64
}

type Resolution struct {
	Width  int
	Height int
}

type ComplexSet struct {
	Resolution Resolution
	Real       Range
	Imaginary  Range
	Iterations int
}

func (c ComplexSet) XStep() float64 {
	return (c.Real.End - c.Real.Start) / float64(c.Resolution.Width)
}

func (c ComplexSet) YStep() float64 {
	return (c.Imaginary.End - c.Imaginary.Start) / float64(c.Resolution.Height)
}

func (conf ComplexSet) Plane(algorithm func(x float64, y float64, iterations int) int) [][]int {
	r := conf.Resolution
	xStep := conf.XStep()
	yStep := conf.YStep()

	result := make([][]int, r.Width)

	for x := 0; x < r.Width; x++ {
		col := make([]int, r.Height)
		for y := 0; y < r.Height; y++ {
			cx := (float64(x) * xStep) + conf.Real.Start
			cy := (float64(y) * yStep) + conf.Imaginary.Start
			count := algorithm(cx, cy, conf.Iterations)
			col[y] = count
		}
		result[x] = col
	}
	return result
}
