package fractal

import (
	"image"
	"image/color"
)

func (p Plane) Image() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, p.grid.Width, p.grid.Height))
	colorCalculator := colorCalculator(float64(p.iterations))
	for x, col := range p.values {
		for y, val := range col {
			color := colorCalculator(float64(val))
			img.Set(x, y, color)
		}
	}
	return img
}

func colorCalculator(iterations float64) func(i float64) color.RGBA {
	colorStep := 255.0 / iterations
	iterationsHalf := iterations / 2.0
	return func(i float64) color.RGBA {
		blue := 0.0
		green := 2.0 * i
		red := 0.0
		if i >= iterationsHalf {
			blue = i - iterationsHalf
		}
		if i > iterationsHalf {
			green = 2.0 * (iterations - i)
		}
		if i <= iterationsHalf {
			red = 255.0 - 2.0*i
		}
		return color.RGBA{
			uint8(red * colorStep),
			uint8(green * colorStep),
			uint8(blue * colorStep),
			0xff}
	}
}
