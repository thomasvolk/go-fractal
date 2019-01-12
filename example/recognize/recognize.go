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

	"github.com/Xamber/Varis"
)

var (
	MAX_ANGLE              = 360.0
	LEARNSET_CONFIG_HEADER = 6
)

func createFile(filename string) *os.File {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return f
}

func openFile(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(num int, rating float64, outputdir string, plane *fractal.Plane,
	angleStep float64, threshold float64) {
	cs := plane.ComplexSet()
	baseFileName := fmt.Sprintf("%03d_Real_%s_Imag_%s_Rating_%f", num, cs.Real, cs.Imaginary, rating)
	fractalFile := createFile(fmt.Sprintf("%s/%s.png", outputdir, baseFileName))
	defer fractalFile.Close()
	shapeImg := image.NewRGBA(image.Rect(0, 0, 100, 100))
	shapeFile := createFile(fmt.Sprintf("%s/%s.shape.png", outputdir, baseFileName))
	defer shapeFile.Close()

	fractImage := plane.ImageWithColorSet("gray")
	cx, cy := plane.Box().Center()
	fractImage.Set(cx, cy, color.RGBA{0, 255, 255, 255})

	shape := plane.Shape(cx, cy, angleStep, threshold)
	for _, p := range shape.Points() {
		fractImage.Set(p.X, p.Y, color.RGBA{255, 0, 0, 255})
	}
	png.Encode(fractalFile, fractImage)

	normalizedShape := shape.Normalize()
	shapeTextFile := createFile(fmt.Sprintf("%s/%s.shape.txt", outputdir, baseFileName))
	defer shapeTextFile.Close()
	shapeTextFile.Write([]byte(fmt.Sprintf("%v\n", rating)))
	for _, value := range normalizedShape {
		x := value[0]
		y := value[1]
		shapeTextFile.Write([]byte(fmt.Sprintf(" %v", x)))
		shapeTextFile.Write([]byte(fmt.Sprintf(" %v", y)))
		shapeImg.Set(
			int(x*float64(shapeImg.Bounds().Dx())),
			int(y*float64(shapeImg.Bounds().Dy())),
			color.RGBA{255, 0, 0, 255})
	}
	png.Encode(shapeFile, shapeImg)

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

func parseIntArray(valuesLine string) []int {
	textValues := strings.Fields(valuesLine)
	values := make([]int, len(textValues))
	for i, v := range textValues {
		values[i] = parseInt(v)
	}
	return values
}

func createLearnSet(sourceFile string) (string, []int) {
	file := openFile(sourceFile)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	width := parseInt(scanner.Text())
	scanner.Scan()
	heigth := parseInt(scanner.Text())
	scanner.Scan()
	outputdir := scanner.Text()
	scanner.Scan()
	threshold := parseFloat(scanner.Text())
	scanner.Scan()
	angleStep := parseFloat(scanner.Text())
	scanner.Scan()
	middleLayer := parseIntArray(scanner.Text())
	scanner.Scan()

	count := LEARNSET_CONFIG_HEADER
	for scanner.Scan() {
		count++
		lineText := scanner.Text()
		if strings.HasPrefix(lineText, "#") {
			continue
		}
		values := strings.Fields(lineText)
		if len(values) != 5 {
			panic(fmt.Sprintf("line %d: wrong file format", count))
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
	return outputdir, middleLayer
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

func readShape(sourceFile string) (varis.Vector, float64) {
	file := openFile(sourceFile)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	rating := parseFloat(scanner.Text())
	scanner.Scan()
	valuesLine := scanner.Text()
	textValues := strings.Fields(valuesLine)
	shapeValues := make(varis.Vector, len(textValues))
	for _, v := range textValues {
		shapeValues = append(shapeValues, parseFloat(v))
	}
	return shapeValues, rating
}

func getLearnSet(dir string) (varis.Dataset, int) {
	files, err := filepath.Glob(fmt.Sprintf("%s/*.shape.txt", dir))
	if err != nil {
		panic(err)
	}
	learnSet := varis.Dataset{}
	shapeSize := 0
	for _, sourceFile := range files {
		shapeValues, rating := readShape(sourceFile)
		shapeSize = len(shapeValues)
		learnSet = append(learnSet, [2]varis.Vector{shapeValues, varis.Vector{rating}})
	}
	return learnSet, shapeSize
}

func learn(learnsetDir string, iterations int, middleLayer []int) varis.Perceptron {
	learnSet, shapeSize := getLearnSet(learnsetDir)
	config := []int{shapeSize}
	config = append(config, middleLayer...)
	config = append(config, 1)
	net := varis.CreatePerceptron(config...)
	trainer := varis.PerceptronTrainer{
		Network: &net,
		Dataset: learnSet,
	}
	trainer.BackPropagation(iterations)
	return net
}

func main() {
	fmt.Println("create learn set ...")
	learnsetDir, middleLayer := createLearnSet("learnset.txt")

	fmt.Println("learn ...")
	net := learn(learnsetDir, 10000, middleLayer)

	fmt.Println("test:")
	learnSet, _ := getLearnSet(learnsetDir)
	for _, entry := range learnSet {
		shapeValues := entry[0]
		expectedRating := entry[1][0]
		result := net.Calculate(shapeValues)
		fmt.Printf("result: %f - expected: %f\n", result[0], expectedRating)
	}

}
