package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	fractal "../.."
)

var (
	MAX_ANGLE              = 360.0
	LEARNSET_CONFIG_HEADER = 6
)

func writeFile(num int, rating float64, outputdir string, plane *fractal.Plane, angleStep float64, threshold float64) {
	cs := plane.ComplexSet()
	baseFileName := fmt.Sprintf("%03d_Real_%s_Imag_%s_Rating_%f", num, cs.Real, cs.Imaginary, rating)
	f, err := os.Create(fmt.Sprintf("%s/%s.png", outputdir, baseFileName))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	shapeImg := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fImg, err := os.Create(fmt.Sprintf("%s/%s.shape.png", outputdir, baseFileName))
	if err != nil {
		panic(err)
	}
	defer fImg.Close()

	fractImage := plane.ImageWithColorSet("gray")
	cx, cy := plane.Box().Center()
	fractImage.Set(cx, cy, color.RGBA{0, 255, 255, 255})

	shape := plane.Shape(cx, cy, angleStep, threshold)
	pointsF := MAX_ANGLE / angleStep
	points := int(pointsF)
	if pointsF > float64(points) {
		points++
	}

	nomrmalzedShape := make([]float64, points*2, points*2)
	nomrmalzedShapeIndex := 0
	for _, p := range shape {
		fractImage.Set(p.X, p.Y, color.RGBA{255, 0, 0, 255})
		normX := float64(p.X) / float64(plane.Width())
		nomrmalzedShape[nomrmalzedShapeIndex] = normX
		nomrmalzedShapeIndex++
		normY := float64(p.Y) / float64(plane.Height())
		nomrmalzedShape[nomrmalzedShapeIndex] = normY
		nomrmalzedShapeIndex++
		shapeImg.Set(
			int(normX*float64(shapeImg.Bounds().Dx())),
			int(normY*float64(shapeImg.Bounds().Dy())),
			color.RGBA{255, 0, 0, 255})
	}
	png.Encode(f, fractImage)
	png.Encode(fImg, shapeImg)

	defer f.Close()
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

func createLearnSet(sourceFile string) string {
	file, err := os.Open(sourceFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	width := parseInt(scanner.Text())
	scanner.Scan()
	heigth := parseInt(scanner.Text())
	scanner.Scan()
	outputdir := scanner.Text()
	scanner.Scan()
	angleStep := parseFloat(scanner.Text())
	scanner.Scan()
	threshold := parseFloat(scanner.Text())
	scanner.Scan()

	count := 0
	for scanner.Scan() {
		count++
		lineText := scanner.Text()
		if strings.HasPrefix(lineText, "#") {
			continue
		}
		values := strings.Fields(lineText)
		if len(values) != 5 {
			panic(fmt.Sprintf("line %d: wrong file format", count+LEARNSET_CONFIG_HEADER))
		}
		x := parseFloat(values[0])
		y := parseFloat(values[1])
		r := parseFloat(values[2])
		iterations := parseInt(values[3])
		rating := parseFloat(values[4])
		p := mandelbrot(width, heigth, x, y, r, r, iterations)
		writeFile(count, rating, outputdir, &p, angleStep, threshold)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return outputdir
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
	learnsetDir := createLearnSet("learnset.txt")
	files, err := filepath.Glob(fmt.Sprintf("%s/*.shape.txt", learnsetDir))
	if err != nil {
		panic(err)
	}
	fmt.Println(files)
}
