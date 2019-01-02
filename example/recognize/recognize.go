package main

import (
	"bufio"
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	fractal "../.."
)

var (
	WIDTH     = 600
	HEIGHT    = 600
	OUTPUTDIR = "learnset"
)

func writeFile(num int, rating float64, outputdir string, plane *fractal.Plane) {
	cs := plane.ComplexSet()
	baseFileName := fmt.Sprintf("%03d_Real_%s_Imag_%s_Rating_%f", num, cs.Real, cs.Imaginary, rating)
	f, err := os.Create(fmt.Sprintf("%s/%s.png", outputdir, baseFileName))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	image := plane.ImageWithColorSet("gray")
	cx, cy := plane.Box().Center()
	image.Set(cx, cy, color.RGBA{0, 255, 255, 255})
	shape := plane.Shape(cx, cy, 1.0, 0.1)

	nomrmalzedShape := make([]float64, 720, 720)
	nomrmalzedShapeIndex := 0
	for _, p := range shape {
		image.Set(p.X, p.Y, color.RGBA{255, 0, 0, 255})
		nomrmalzedShape[nomrmalzedShapeIndex] = float64(p.X) / float64(plane.Width())
		nomrmalzedShapeIndex++
		nomrmalzedShape[nomrmalzedShapeIndex] = float64(p.Y) / float64(plane.Height())
		nomrmalzedShapeIndex++
	}
	png.Encode(f, image)

	f, err = os.Create(fmt.Sprintf("%s/%s.shape.txt", outputdir, baseFileName))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write([]byte(fmt.Sprintf("%v", rating)))
	for _, value := range nomrmalzedShape {
		f.Write([]byte(fmt.Sprintf(" %v", value)))
	}
}

func parseFloat(value string) float64 {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}
	return v
}

func parseInt(value string) int {
	v, err := strconv.Atoi(value)
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
		lineText := scanner.Text()
		if strings.HasPrefix(lineText, "#") {
			continue
		}
		values := strings.Fields(lineText)
		if len(values) != 5 {
			panic(fmt.Sprintf("line %d: wrong file format", line))
		}
		x := parseFloat(values[0])
		y := parseFloat(values[1])
		r := parseFloat(values[2])
		iterations := parseInt(values[3])
		rating := parseFloat(values[4])
		p := mandelbrot(WIDTH, HEIGHT, x, y, r, r, iterations)
		writeFile(line, rating, OUTPUTDIR, &p)
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

	return m.Plane(width, height, iterations)
}

func main() {
	createLearnSet("learnset.txt")
	files, err := filepath.Glob(fmt.Sprintf("%s/*.shape.txt", OUTPUTDIR))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)
}
