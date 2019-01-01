package fractal

import (
	"reflect"
	"testing"
)

func TestValue(t *testing.T) {
	iterations := 65
	c := ComplexSet{
		Real:      Range{-1.5, 1.5},
		Imaginary: Range{-1.0, 1.0},
		Algorithm: Mandelbrot,
	}
	p := c.Plane(300, 200, iterations)
	cx, cy := p.Box().Center()
	if cx != 150 || cy != 100 {
		t.Errorf("wrong center: x=%d, y=%d", cx, cy)
	}
	v := p.Value(cx, cy)
	expected := iterations
	if v != expected {
		t.Errorf("wrong value: %d != %d", expected, v)
	}
	v = p.Value(0, 0)
	expected = 1
	if v != expected {
		t.Errorf("wrong value: %d != %d", expected, v)
	}
}

func TestMean(t *testing.T) {
	src := [][]int{[]int{1, 2, 3, 6}, []int{1, 2, 3, 6}, []int{1, 2, 0, 3}}
	expected := 2.5
	m := mean(src)
	if expected != m {
		t.Errorf("wrong mean: %f != %f", expected, m)
	}
}

func TestDeviationZero(t *testing.T) {
	zero := [][]int{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}, []int{1, 1, 1, 1}}
	expected := 0.0
	d := deviation(zero)
	if expected != d {
		t.Errorf("wrong deviation: %f != %f", expected, d)
	}
}

func TestDeviation(t *testing.T) {
	zero := [][]int{[]int{0, 0, 0}, []int{0, 0, 9}, []int{0, 0, 0}}
	expected := 2.8284271247461903
	d := deviation(zero)
	if expected != d {
		t.Errorf("wrong deviation: %f != %f", expected, d)
	}
}

func TestCrop(t *testing.T) {
	src := [][]int{[]int{1, 2, 3, 4},
		[]int{10, 20, 30, 40},
		[]int{100, 200, 300, 400},
		[]int{1000, 2000, 3000, 4000}}
	part := crop(src, Box{1, 1, 2, 2})
	expected := [][]int{[]int{20, 30}, []int{200, 300}}
	if !reflect.DeepEqual(part, expected) {
		t.Errorf("part has not the expected layout: %v != %v", expected, part)
	}
}

func TestXStepYStep(t *testing.T) {
	c := ComplexSet{
		Real:      Range{-1.5, 1.5},
		Imaginary: Range{-1.0, 1.0},
		Algorithm: Mandelbrot,
	}
	p := c.Plane(300, 200, 10)
	expected := 0.01
	result := p.XStep()
	if expected != result {
		t.Errorf("Xstep %v != %v", expected, result)
	}
	expected = 0.01
	result = p.YStep()
	if expected != result {
		t.Errorf("Ystep %v != %v", expected, result)
	}
}

func TestPlaneCrop(t *testing.T) {
	c := ComplexSet{
		Real:      Range{-1.5, 1.5},
		Imaginary: Range{-1.0, 1.0},
		Algorithm: Mandelbrot,
	}
	p := c.Plane(300, 200, 10)
	pc := p.Crop(Box{0, 0, 150, 100})

	expected := 0.01
	result := pc.XStep()
	if expected != result {
		t.Errorf("Xstep %v != %v", expected, result)
	}
	expected = 0.01
	result = pc.YStep()
	if expected != result {
		t.Errorf("Ystep %v != %v", expected, result)
	}
	cc := pc.ComplexSet()
	if cc.Real.Start != -1.5 || cc.Real.End != 0.0 || cc.Imaginary.Start != -1.0 || cc.Imaginary.End != 0.0 {
		t.Errorf("ComplexSet")
	}

}
