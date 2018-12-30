package fractal

import "fmt"

type Box struct {
	X      int
	Y      int
	Width  int
	Height int
}

func (b Box) String() string {
	return fmt.Sprintf("Box(x=%d, y=%d, w=%d, h=%d)", b.X, b.Y, b.Width, b.Height)
}

func (outer Box) InnerBox(x int, y int) Box {
	nx := (x - outer.X)
	ny := (y - outer.Y)
	xmin := (outer.Width - nx)
	ymin := (outer.Height - ny)
	if nx < xmin {
		xmin = nx
	}
	if ny < ymin {
		ymin = ny
	}

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
