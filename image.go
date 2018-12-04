package fractal

import (
	"image"
	"image/color"
)

func Image(conf Config, s IterationMapping) *image.RGBA {
	r := conf.Resolution
	img := image.NewRGBA(image.Rect(0, 0, r.Width, r.Height))
	colorCalculator := colorCalculator(conf)
	for y := 0; y < r.Height; y++ {
		for x := 0; x < r.Width; x++ {
			color := colorCalculator(float64(s.Get(x, y)))
			img.Set(x, y, color)
		}
	}
	return img
}

func colorCalculator(conf Config) func(i float64) color.RGBA {
	iterations := float64(conf.Iterations)
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
