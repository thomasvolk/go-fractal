
all: build
	
build: test
	go build mandelbrot-cli.go

test:
	(cd .. && go test)

clean:
	rm -f mandelbrot-cli *.png

test-run: build
	./mandelbrot-cli --zoom 50 --iterations 600
