
all: build
	
build: test
	go build mandelbrot-cli.go

test:
	(cd ../.. && go test ./...)

clean:
	rm -rf mandelbrot-cli *.png result

test-run: build
	./mandelbrot-cli --zoom 20 --iterations 600 --outputdir result/

