package fractal

type Range struct {
	Start float64
	End   float64
}

type Resolution struct {
	Width  int
	Height int
}

type Config struct {
	Resolution Resolution
	Real       Range
	Imaginary  Range
	Iterations int
}

func (c Config) XStep() float64 {
	return (c.Real.End - c.Real.Start) / float64(c.Resolution.Width)
}

func (c Config) YStep() float64 {
	return (c.Imaginary.End - c.Imaginary.Start) / float64(c.Resolution.Height)
}
