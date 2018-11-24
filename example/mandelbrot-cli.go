package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"net/http"
	"os"
	"strconv"

	mandelbrot ".."
)

func writeFile(outputfile string, image *image.RGBA) {
	f, err := os.Create(outputfile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, image)
}

func forQueryParam(r *http.Request, param string, f func(value float64)) {
	values, ok := r.URL.Query()[param]
	if ok {
		fval, err := strconv.ParseFloat(values[0], 64)
		if err == nil {
			f(fval)
		}
	}
}

func drawHandler(m mandelbrot.Mandelbrot) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		mc := mandelbrot.Mandelbrot{
			Xstart:     m.Xstart,
			Xend:       m.Xend,
			Ystart:     m.Ystart,
			Yend:       m.Yend,
			Iterations: m.Iterations,
			Width:      m.Width,
			Height:     m.Height,
		}
		w.Header().Set("Content-Type", "image/png")
		forQueryParam(r, "xstart", func(value float64) { mc.Xstart = value })
		forQueryParam(r, "xend", func(value float64) { mc.Xend = value })
		forQueryParam(r, "ystart", func(value float64) { mc.Ystart = value })
		forQueryParam(r, "yend", func(value float64) { mc.Yend = value })
		forQueryParam(r, "iterations", func(value float64) { mc.Iterations = int(value) })
		image := mc.Draw()
		png.Encode(w, image)
	}
}

func main() {
	var xstart float64
	var xend float64
	var ystart float64
	var yend float64
	var iterations int
	var width int
	var height int
	var outputfile string
	var serve bool
	var port int

	flag.Float64Var(&xstart, "xstart", -2.0, "xstart")
	flag.Float64Var(&xend, "xend", 1.2, "xend")
	flag.Float64Var(&ystart, "ystart", -1.2, "ystart")
	flag.Float64Var(&yend, "yend", 1.2, "yend")
	flag.IntVar(&iterations, "iterations", 100, "iterations")
	flag.IntVar(&width, "width", 400, "width")
	flag.IntVar(&height, "height", 300, "height")
	flag.StringVar(&outputfile, "outputfile", "mandelbrot.png", "outputfile")
	flag.BoolVar(&serve, "serve", false, "start http server")
	flag.IntVar(&port, "port", 8080, "http port")

	flag.Parse()

	m := mandelbrot.Mandelbrot{
		Xstart:     xstart,
		Xend:       xend,
		Ystart:     ystart,
		Yend:       yend,
		Iterations: iterations,
		Width:      width,
		Height:     height,
	}

	if serve {
		http.HandleFunc("/", drawHandler(m))
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
			panic(err)
		}
	} else {
		image := m.Draw()
		writeFile(outputfile, image)
	}
}
