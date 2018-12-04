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
	XRange     Range
	YRange     Range
	Iterations int
}

type IterationMapping struct {
	mapping [][]int
}

func NewIterationMapping(mapping [][]int) IterationMapping { return IterationMapping{mapping} }
func (im IterationMapping) Get(x int, y int) int           { return im.mapping[y][x] }
