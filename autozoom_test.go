package fractal

import (
	"reflect"
	"testing"
)

func TestPart(t *testing.T) {
	src := [][]int{[]int{1, 2, 3, 4},
		[]int{10, 20, 30, 40},
		[]int{100, 200, 300, 400},
		[]int{1000, 2000, 3000, 4000}}
	part := part(src, Box{1, 1, 2, 2})
	expected := [][]int{[]int{20, 30}, []int{200, 300}}
	if !reflect.DeepEqual(part, expected) {
		t.Errorf("part has not the expected layout: %v != %v", expected, part)
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

func TestInnerBox(t *testing.T) {
	box := Box{0, 0, 100, 100}
	ib := innerBox(0, 0, box)
	if (ib != Box{0, 0, 0, 0}) {
		t.Errorf("innerBox result failure %s", ib)
	}
	ib = innerBox(99, 99, box)
	if (ib != Box{98, 98, 2, 2}) {
		t.Errorf("innerBox result failure %s", ib)
	}
	ib = innerBox(50, 50, box)
	if (ib != Box{0, 0, 100, 100}) {
		t.Errorf("innerBox result failure %s", ib)
	}
	ib = innerBox(25, 75, box)
	if (ib != Box{0, 50, 50, 50}) {
		t.Errorf("innerBox result failure %s", ib)
	}
	ib = innerBox(2, 75, box)
	if (ib != Box{0, 73, 4, 4}) {
		t.Errorf("innerBox result failure %s", ib)
	}
	ib = innerBox(60, 60, Box{10, 10, 100, 100})
	if (ib != Box{10, 10, 100, 100}) {
		t.Errorf("innerBox result failure %s", ib)
	}
}
