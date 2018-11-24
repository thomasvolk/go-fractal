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
	var wg sync.WaitGroup
	xStep := (m.Xend - m.Xstart) / float64(m.Width)
	yStep := (m.Yend - m.Ystart) / float64(m.Height)
	wg.Add(m.Width * m.Height)
	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			go func(image *image.RGBA, px int, py int) {
				defer wg.Done()
				m.drawPoint(image, px, py, xStep, yStep)
			}(img, x, y)
		}
	}
	wg.Wait()
	return img
}

func (m *Mandelbrot) drawPoint(img *image.RGBA, x int, y int, xStep float64, yStep float64) {
	color := color.RGBA{0x00, 0x00, 0x00, 0xff}
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
	color = m.calculateColor(i)
	img.Set(int(x), int(y), color)
}

func (m *Mandelbrot) calculateColor(i int) color.RGBA {
	colorStep := 255.0 / float64(m.Iterations)
	blue := 0.0
	green := 2.0 * float64(i)
	red := 0.0
	if i >= m.Iterations/2 {
		blue = (float64(i) - float64(m.Iterations)/2.0)
	}
	if i > m.Iterations/2 {
		green = float64(2 * (m.Iterations - i))
	}
	if i <= m.Iterations/2 {
		red = 255.0 - 2.0*float64(i)
	}
	return color.RGBA{
		uint8(red * colorStep),
		uint8(green * colorStep),
		uint8(blue * colorStep),
		0xff}
}
