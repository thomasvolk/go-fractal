package fractal

import (
	"testing"
)

func TestAbs(t *testing.T) {
	test := func(expected, result int) {
		if expected != result {
			t.Errorf("abs: %d != %d", expected, result)
		}
	}

	test(0, abs(0))
	test(0, abs(-0))
	test(1, abs(-1))
	test(1, abs(1))
	test(99, abs(99))
	test(8, abs(-8))
}

func TestMaxPitch(t *testing.T) {
	test := func(expected, result Value) {
		if expected != result {
			t.Errorf("abs: %v != %v", expected, result)
		}
	}

	test(Value{0, 0, 10}, maxPitch([]Value{
		Value{0, 0, 10},
	}, 1, 0.1))
	test(Value{2, 0, 100}, maxPitch([]Value{
		Value{0, 0, 20},
		Value{1, 0, 50},
		Value{2, 0, 100},
		Value{3, 0, 110},
	}, 120, 0.4))
	test(Value{0, 0, 10}, maxPitch([]Value{
		Value{0, 0, 10},
		Value{1, 0, 10},
		Value{2, 0, 10},
		Value{3, 0, 10},
		Value{4, 0, 10},
	}, 300, 0.5))
	test(Value{2, 0, -100}, maxPitch([]Value{
		Value{0, 0, 20},
		Value{1, 0, 50},
		Value{2, 0, -100},
		Value{3, 0, -110},
	}, 120, 0.4))
}

func TestCreateShape(t *testing.T) {
	iterations := 10
	c := ComplexSet{
		Real:      Range{-1.5, 1.5},
		Imaginary: Range{-1.0, 1.0},
		Algorithm: Mandelbrot,
	}
	p := c.Plane(400, 300, iterations)
	frames := p.RasterFrames(3)
	f := frames[3]
	cx, cy := f.Box().Center()
	f.Shape(cx, cy, 20, 0.03)
}
