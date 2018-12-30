package fractal

import (
	"fmt"
	"image"
	"image/color"
)

var (
	COLOR_SET_FACTORY = map[string]func(iterations float64) func(i float64) color.RGBA{
		"default": defaultColorSetFactory,
		"gray":    grayColorSetFactory,
	}
)

func (p Plane) Image() *image.RGBA {
	return p.ImageWithColorSet("default")
}

func (p Plane) ImageWithColorSet(colorSet string) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, p.width, p.height))
	colorCalculator := getColorSet(colorSet, float64(p.iterations))
	for x, col := range p.values {
		for y, val := range col {
			color := colorCalculator(float64(val))
			img.Set(x, y, color)
		}
	}
	return img
}

func getColorSet(colorSet string, iterations float64) func(i float64) color.RGBA {
	colorSetFactory := COLOR_SET_FACTORY[colorSet]
	if colorSetFactory == nil {
		panic(fmt.Sprintf("unknown color set: %s", colorSet))
	}
	return colorSetFactory(iterations)
}

func grayColorSetFactory(iterations float64) func(i float64) color.RGBA {
	colorStep := 255.0 / iterations
	return func(i float64) color.RGBA {
		allColors := 255.0 - colorStep*i
		return color.RGBA{
			uint8(allColors),
			uint8(allColors),
			uint8(allColors),
			0xff}
	}
}

func defaultColorSetFactory(iterations float64) func(i float64) color.RGBA {
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
