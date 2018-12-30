package fractal

import (
	"testing"

	model "./model"
)

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

func TestCrop(t *testing.T) {
	c := ComplexSet{
		Real:      Range{-1.5, 1.5},
		Imaginary: Range{-1.0, 1.0},
		Algorithm: Mandelbrot,
	}
	p := c.Plane(300, 200, 10)
	pc := p.Crop(model.Box{0, 0, 150, 100})

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
