package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"

	fractal "../.."
)

func writeFile(num int, outputdir string, plane *fractal.Plane, colorSet string) {
	cs := plane.ComplexSet()
	f, err := os.Create(fmt.Sprintf("%s/%03d_Real_%s_Imag_%s.png", outputdir, num, cs.Real, cs.Imaginary))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, plane.ImageWithColorSet(colorSet))
}

func main() {
	var x float64
	var xradius float64
	var y float64
	var yradius float64
	var iterations int
	var width int
	var height int
	var outputdir string
	var zoom int
	var zoomFactor float64
	var colorSet string
	var slideStep int

	flag.Float64Var(&x, "x", -0.6, "xstart")
	flag.Float64Var(&xradius, "xradius", 1.6, "xradius")
	flag.Float64Var(&y, "y", 0.0, "ystart")
	flag.Float64Var(&yradius, "yradius", 1.2, "yradius")
	flag.IntVar(&iterations, "iterations", 100, "iterations")
	flag.IntVar(&width, "width", 400, "width")
	flag.IntVar(&height, "height", 300, "height")
	flag.StringVar(&outputdir, "outputdir", ".", "outputdir")
	flag.IntVar(&zoom, "zoom", 0, "zoom")
	flag.Float64Var(&zoomFactor, "zoom-factor", 0.5, "zoom factor valid value 1 > and > 0 ")
	flag.StringVar(&colorSet, "color-set", "default", "colorset [ default | gray ] default is 'default'")
	flag.IntVar(&slideStep, "slide-step", 10, "sliding step for zoom")

	flag.Parse()

	m := fractal.ComplexSet{
		Real:      fractal.NewRange(x, xradius),
		Imaginary: fractal.NewRange(y, yradius),
		Algorithm: fractal.Mandelbrot,
	}

	os.MkdirAll(outputdir, os.ModePerm)
	p := m.Plane(width, height, iterations)
	writeFile(0, outputdir, &p, colorSet)

	for z := 0; z < zoom; z++ {
		p = p.AutoZoom(zoomFactor, slideStep)
		writeFile(z+1, outputdir, &p, colorSet)
	}
}
