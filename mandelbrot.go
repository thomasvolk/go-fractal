package mandelbrot

import (
	"image"
	"image/color"
	"math/cmplx"
	"sync"
)

type Mandelbrot struct {
	Xstart     float64
	Xend       float64
	Ystart     float64
	Yend       float64
	Iterations int
	Width      int
	Height     int
}

func (m *Mandelbrot) Draw() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, m.Width, m.Height))
	xStep := (m.Xend - m.Xstart) / float64(m.Width)
	yStep := (m.Yend - m.Ystart) / float64(m.Height)

	var wg sync.WaitGroup
	wg.Add(m.Width * m.Height)

	colorCalculator := m.colorCalculator()
	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			go func(image *image.RGBA, px int, py int) {
				defer wg.Done()
				i := m.calculateIterations(px, py, xStep, yStep)
				color := colorCalculator(float64(i))
				image.Set(int(px), int(py), color)
			}(img, x, y)
		}
	}

	wg.Wait()
	return img
}

func (m *Mandelbrot) calculateIterations(x int, y int, xStep float64, yStep float64) int {
	c := complex((float64(x)*xStep)+m.Xstart,
		(float64(y)*yStep)+m.Ystart)
	z := 0 + 0i
	i := 0
	for ; i < m.Iterations; i++ {
		z = z*z + c
		if cmplx.Abs(z) >= 2 {
			break
		}
	}
	return i
}

func (m *Mandelbrot) colorCalculator() func(i float64) color.RGBA {
	iterations := float64(m.Iterations)
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
