package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	fractal "../.."

	"github.com/Xamber/Varis"
	_ "github.com/Xamber/Varis"
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

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func writeFile(num int, cs *fractal.ComplexSet, width, height, iterations int,
	rating float64, outputdir string,
	shapeSize int, threshold float64) {
	plane := cs.Plane(width, height, iterations)
	baseFileName := fmt.Sprintf("%03d_Real_%s_Imag_%s_Rating_%f", num, cs.Real, cs.Imaginary, rating)
	fractalFile := createFile(fmt.Sprintf("%s/%s.png", outputdir, baseFileName))
	defer fractalFile.Close()
	shapeImg := image.NewRGBA(image.Rect(0, 0, 100, 100))
	shapeFile := createFile(fmt.Sprintf("%s/%s.shape.png", outputdir, baseFileName))
	defer shapeFile.Close()

	fractImage := plane.ImageWithColorSet("gray")
	cx, cy := plane.Box().Center()
	fractImage.Set(cx, cy, color.RGBA{0, 255, 255, 255})

	shape := plane.Shape(cx, cy, shapeSize, threshold)
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
		shapeTextFile.Write([]byte(fmt.Sprintf("%v ", x)))
		shapeTextFile.Write([]byte(fmt.Sprintf("%v ", y)))
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

func createLearnSet(sourceFile, outputdir string, threshold float64,
	shapeSize, width, height int) {
	os.Mkdir(outputdir, os.ModePerm)
	file := openFile(sourceFile)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
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
		m := fractal.ComplexSet{
			Real:      fractal.NewRange(x, r),
			Imaginary: fractal.NewRange(y, r),
			Algorithm: fractal.Mandelbrot,
		}
		writeFile(count, &m, width, height, iterations,
			rating, outputdir, shapeSize, threshold)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
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
	for i, v := range textValues {
		shapeValues[i] = parseFloat(v)
	}
	return shapeValues, rating
}

func getLearnSet(dir string) varis.Dataset {
	files, err := filepath.Glob(fmt.Sprintf("%s/*.shape.txt", dir))
	if err != nil {
		panic(err)
	}
	learnSet := varis.Dataset{}
	for _, sourceFile := range files {
		shapeValues, rating := readShape(sourceFile)
		fmt.Printf("load shape %s - len: %d\n", sourceFile, len(shapeValues))
		learnSet = append(learnSet, [2]varis.Vector{shapeValues, varis.Vector{rating}})
	}
	return learnSet
}

func learn(learnsetDir string, inputLayer, iterations int, middleLayer []int) varis.Perceptron {
	learnSet := getLearnSet(learnsetDir)
	config := []int{inputLayer}
	config = append(config, middleLayer...)
	config = append(config, 1)
	fmt.Printf("create net %v\n", config)
	net := varis.CreatePerceptron(config...)
	trainer := varis.PerceptronTrainer{
		Network: &net,
		Dataset: learnSet,
	}
	trainer.BackPropagation(iterations)
	return net
}

func main() {
	var x float64
	var xradius float64
	var y float64
	var yradius float64
	var iterations int
	var width int
	var height int
	var zoomoutputdir string
	var learnSetDir, learnSetFile string
	var shapeThreshold float64
	var shapeSize int
	var learnIterations int
	var middleLayer string
	var netFile string

	flag.StringVar(&learnSetDir, "learnset", "learnset", "learn set result dir")
	flag.StringVar(&learnSetFile, "learnset-source", "learnset.txt", "learn set source file")
	flag.Float64Var(&x, "x", -0.6, "xstart")
	flag.Float64Var(&xradius, "xradius", 1.6, "xradius")
	flag.Float64Var(&y, "y", 0.0, "ystart")
	flag.Float64Var(&yradius, "yradius", 1.2, "yradius")
	flag.IntVar(&iterations, "iterations", 100, "iterations")
	flag.IntVar(&width, "width", 400, "width")
	flag.IntVar(&height, "height", 300, "height")
	flag.StringVar(&zoomoutputdir, "zoomout", "toom", "zoom outputdir")
	flag.IntVar(&shapeSize, "shape-size", 9, "count of shape points")
	flag.Float64Var(&shapeThreshold, "shape-threshold", 0.03, "threshold for detectiong the shape")
	flag.IntVar(&learnIterations, "learn", 6000, "count of learn steps")
	flag.StringVar(&middleLayer, "middle-layer", "19", "layout of the neuron middle layer")
	flag.StringVar(&netFile, "net", "net.json", "net output file")

	flag.Parse()

	if !exists(learnSetDir) {
		fmt.Println("# create learn set ...")
		createLearnSet(learnSetFile, learnSetDir, shapeThreshold, shapeSize, width, height)
	}

	var net varis.Perceptron

	if !exists(netFile) {
		fmt.Println("# learn ...")
		net = learn(learnSetDir, shapeSize*2, learnIterations, parseIntArray(middleLayer))
		netJson := varis.ToJSON(net)
		file := createFile(netFile)
		defer file.Close()
		file.Write([]byte(netJson))
	} else {
		netJson, err := ioutil.ReadFile(netFile)

		if err != nil {
			panic(err)
		}

		net = varis.FromJSON(string(netJson))
	}

	fmt.Println("# test:")
	learnSet := getLearnSet(learnSetDir)
	for _, entry := range learnSet {
		shapeValues := entry[0]
		expectedRating := entry[1][0]
		result := net.Calculate(shapeValues)
		fmt.Printf("result: %f - expected: %f\n", result[0], expectedRating)
	}

	fmt.Println("# zoom:")
	m := fractal.ComplexSet{
		Real:      fractal.NewRange(x, xradius),
		Imaginary: fractal.NewRange(y, yradius),
		Algorithm: fractal.Mandelbrot,
	}
	p := m.Plane(width, height, iterations)
	division := 2
	frames := p.RasterFrames(division)
	for _, f := range frames {
		cx, cy := f.Box().Center()
		s := f.Shape(cx, cy, shapeSize, shapeThreshold)
		sn := s.Normalize()
		shapeValues := make([]float64, len(sn)*2)
		count := 0
		for _, p := range sn {
			shapeValues[count] = p[0]
			count++
			shapeValues[count] = p[1]
			count++
		}
		result := net.Calculate(shapeValues)
		fmt.Println(result)
	}
}
