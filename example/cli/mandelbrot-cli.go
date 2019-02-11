package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"
	"strconv"
	"strings"

	fractal "../.."
)

type Zoomer interface {
	Zoom(fractal.Plane) fractal.Plane
}

type RasterAutoZoom struct {
	division int
}

type CircleAutoZoom struct {
	x             int
	y             int
	radiusDivisor int
	angleStep     int
}

func (r RasterAutoZoom) Zoom(p fractal.Plane) fractal.Plane {
	return p.RasterAutoZoom(r.division)
}

func (c CircleAutoZoom) Zoom(p fractal.Plane) fractal.Plane {
	newPlane := p.CircleAutoZoom(c.x, c.y, float64(c.radiusDivisor), float64(c.angleStep))
	c.x = newPlane.Width() / 2
	c.y = newPlane.Height() / 2
	return newPlane
}

func getZoomer(zoomConfig string) Zoomer {
	zoom := strings.Split(zoomConfig, ":")
	zoomType := zoom[0]
	if zoomType == "raster" {
		division, err := strconv.Atoi(zoom[1])
		if err != nil {
			panic(err)
		}
		return RasterAutoZoom{division}
	}
	if zoomType == "circle" {
		x, err := strconv.Atoi(zoom[1])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(zoom[2])
		if err != nil {
			panic(err)
		}
		radiusDivisor := 5
		if len(zoom) > 3 {
			radiusDivisor, err = strconv.Atoi(zoom[3])
			if err != nil {
				panic(err)
			}
		}
		angleStep := 6
		if len(zoom) > 4 {
			angleStep, err = strconv.Atoi(zoom[4])
			if err != nil {
				panic(err)
			}
		}
		return CircleAutoZoom{x, y, radiusDivisor, angleStep}
	}
	panic(fmt.Sprintf("unknown zoom type: %s\n", zoomType))
}

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
	var zoomConfig string
	var colorSet string

	flag.Float64Var(&x, "x", -0.6, "xstart")
	flag.Float64Var(&xradius, "xradius", 1.6, "xradius")
	flag.Float64Var(&y, "y", 0.0, "ystart")
	flag.Float64Var(&yradius, "yradius", 1.2, "yradius")
	flag.IntVar(&iterations, "iterations", 100, "iterations")
	flag.IntVar(&width, "width", 400, "width")
	flag.IntVar(&height, "height", 300, "height")
	flag.StringVar(&outputdir, "outputdir", ".", "outputdir")
	flag.IntVar(&zoom, "zoom", 0, "zoom")
	flag.StringVar(&zoomConfig, "zoom-config", "raster:2", "zoom type [ raster:N | circle:x:y:R:S ] default is 'raster:2'")
	flag.StringVar(&colorSet, "color-set", "default", "colorset [ default | gray ] default is 'default'")

	flag.Parse()

	m := fractal.ComplexSet{
		Real:      fractal.NewRange(x, xradius),
		Imaginary: fractal.NewRange(y, yradius),
		Algorithm: fractal.Mandelbrot,
	}

	os.MkdirAll(outputdir, os.ModePerm)
	p := m.Plane(width, height, iterations)
	writeFile(0, outputdir, &p, colorSet)

	zoomer := getZoomer(zoomConfig)
	for z := 0; z < zoom; z++ {
		p = zoomer.Zoom(p)
		p = p.Scale(width, height)
		writeFile(z+1, outputdir, &p, colorSet)
	}
}
