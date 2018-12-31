package fractal

import (
	"fmt"
)

type Box struct {
	X      int
	Y      int
	Width  int
	Height int
}

func (b Box) String() string {
	return fmt.Sprintf("Box(x=%d, y=%d, w=%d, h=%d)", b.X, b.Y, b.Width, b.Height)
}

func (b Box) Center() (int, int) {
	return b.X + b.Width/2, b.Y + b.Height/2
}

func (outer Box) InnerBox(x int, y int) Box {
	left := (x - outer.X)
	up := (y - outer.Y)
	right := (outer.Width - left)
	down := (outer.Height - up)

	xmin := min(left, right)
	ymin := min(up, down)

	widthScaleFactor := float64(xmin*2) / float64(outer.Height)
	heightScaleFactor := float64(ymin*2) / float64(outer.Height)

	scaleFactor := heightScaleFactor
	if widthScaleFactor < heightScaleFactor {
		scaleFactor = widthScaleFactor
	}

	newWidth := int(float64(outer.Width) * scaleFactor)
	newHeight := int(float64(outer.Height) * scaleFactor)

	return Box{x - newWidth/2, y - newHeight/2, newWidth, newHeight}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
