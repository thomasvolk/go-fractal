package mandelbrot

import (
	"image"
	"image/color"
)

func (ms MandelbrotSet) Image() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, ms.Width, ms.Height))
	colorCalculator := ms.colorCalculator()
	for y := range ms.Set {
		for x := range ms.Set[y] {
			color := colorCalculator(float64(ms.Set[y][x]))
			img.Set(x, y, color)
		}
	}
	return img
}

func (ms MandelbrotSet) colorCalculator() func(i float64) color.RGBA {
	iterations := float64(ms.Iterations)
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
