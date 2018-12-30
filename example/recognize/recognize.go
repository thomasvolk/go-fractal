package main

import (
	"bufio"
	"fmt"
	"image/png"
	"os"
	"strconv"
	"strings"

	fractal "../.."
)

var (
	WIDTH      = 600
	HEIGHT     = 600
	ITERATIONS = 180
	OUTPUTDIR  = "learnset"
)

func writeFile(num int, outputdir string, plane *fractal.Plane) {
	cs := plane.ComplexSet()
	f, err := os.Create(fmt.Sprintf("%s/%03d_Real_%s_Imag_%s.png", outputdir, num, cs.Real, cs.Imaginary))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, plane.ImageWithColorSet("gray"))
}

func parse(value string) float64 {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}
	return v
}

func createLearnSet(sourceFile string) {
	file, err := os.Open(sourceFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		line++
		values := strings.Fields(scanner.Text())
		if len(values) != 3 {
			panic(fmt.Sprintf("line %d: wrong file format", line))
		}
		x := parse(values[0])
		y := parse(values[1])
		r := parse(values[2])
		p := mandelbrot(WIDTH, HEIGHT, x, y, r, r, ITERATIONS)
		writeFile(line, OUTPUTDIR, &p)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func mandelbrot(width int, height int, x float64, y float64, xradius float64, yradius float64,
	iterations int) fractal.Plane {
	m := fractal.ComplexSet{
		Real:      fractal.NewRange(x, xradius),
		Imaginary: fractal.NewRange(y, yradius),
		Algorithm: fractal.Mandelbrot,
	}

	return fractal.NewPlane(m, width, height, iterations)
}

func main() {
	createLearnSet("learnset.txt")
}
