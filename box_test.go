package fractal

import (
	"testing"
)

func TestInnerBox(t *testing.T) {
	box := Box{0, 0, 100, 100}
	ib := box.InnerBox(0, 0)
	if (ib != Box{0, 0, 0, 0}) {
		t.Errorf("innerBox result failure %s", ib)
	}
	ib = box.InnerBox(99, 99)
	if (ib != Box{98, 98, 2, 2}) {
		t.Errorf("innerBox result failure %s", ib)
	}
	ib = box.InnerBox(50, 50)
	if (ib != Box{0, 0, 100, 100}) {
		t.Errorf("innerBox result failure %s", ib)
	}
	ib = box.InnerBox(25, 75)
	if (ib != Box{0, 50, 50, 50}) {
		t.Errorf("innerBox result failure %s", ib)
	}
	ib = box.InnerBox(2, 75)
	if (ib != Box{0, 73, 4, 4}) {
		t.Errorf("innerBox result failure %s", ib)
	}
	ib = Box{10, 10, 100, 100}.InnerBox(60, 60)
	if (ib != Box{10, 10, 100, 100}) {
		t.Errorf("innerBox result failure %s", ib)
	}
}
