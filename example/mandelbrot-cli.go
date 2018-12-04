package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"net/http"
	"os"
	"strconv"

	fractal ".."
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

func drawHandler(conf fractal.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		copyConf := fractal.Config{
			Real:       fractal.Range{Start: conf.Real.Start, End: conf.Real.End},
			Imaginary:  fractal.Range{Start: conf.Imaginary.Start, End: conf.Imaginary.End},
			Iterations: conf.Iterations,
			Resolution: fractal.Resolution{Width: conf.Resolution.Width, Height: conf.Resolution.Height},
		}
		w.Header().Set("Content-Type", "image/png")
		forQueryParam(r, "xstart", func(value float64) { copyConf.Real.Start = value })
		forQueryParam(r, "xend", func(value float64) { copyConf.Real.End = value })
		forQueryParam(r, "ystart", func(value float64) { copyConf.Imaginary.Start = value })
		forQueryParam(r, "yend", func(value float64) { copyConf.Imaginary.End = value })
		forQueryParam(r, "iterations", func(value float64) { copyConf.Iterations = int(value) })
		image := copyConf.Image(fractal.Mandelbrot)
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

	conf := fractal.Config{
		Real:       fractal.Range{Start: xstart, End: xend},
		Imaginary:  fractal.Range{Start: ystart, End: yend},
		Iterations: iterations,
		Resolution: fractal.Resolution{Width: width, Height: height},
	}

	if serve {
		http.HandleFunc("/", drawHandler(conf))
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
			panic(err)
		}
	} else {
		image := conf.Image(fractal.Mandelbrot)
		writeFile(outputfile, image)
	}
}
