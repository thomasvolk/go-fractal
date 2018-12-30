package utils

import (
	"reflect"
	"testing"

	model "../model"
)

func TestMean(t *testing.T) {
	src := [][]int{[]int{1, 2, 3, 6}, []int{1, 2, 3, 6}, []int{1, 2, 0, 3}}
	expected := 2.5
	m := Mean(src)
	if expected != m {
		t.Errorf("wrong mean: %f != %f", expected, m)
	}
}

func TestDeviationZero(t *testing.T) {
	zero := [][]int{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}, []int{1, 1, 1, 1}}
	expected := 0.0
	d := Deviation(zero)
	if expected != d {
		t.Errorf("wrong deviation: %f != %f", expected, d)
	}
}

func TestDeviation(t *testing.T) {
	zero := [][]int{[]int{0, 0, 0}, []int{0, 0, 9}, []int{0, 0, 0}}
	expected := 2.8284271247461903
	d := Deviation(zero)
	if expected != d {
		t.Errorf("wrong deviation: %f != %f", expected, d)
	}
}

func TestCrop(t *testing.T) {
	src := [][]int{[]int{1, 2, 3, 4},
		[]int{10, 20, 30, 40},
		[]int{100, 200, 300, 400},
		[]int{1000, 2000, 3000, 4000}}
	part := Crop(src, model.Box{1, 1, 2, 2})
	expected := [][]int{[]int{20, 30}, []int{200, 300}}
	if !reflect.DeepEqual(part, expected) {
		t.Errorf("part has not the expected layout: %v != %v", expected, part)
	}
}
