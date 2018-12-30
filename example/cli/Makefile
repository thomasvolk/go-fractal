
all: build
	
build: test
	go build mandelbrot-cli.go

test:
	(cd ../.. && go test ./...)

clean:
	rm -rf mandelbrot-cli *.png result

test-run: build
	mkdir -p result/default/
	./mandelbrot-cli --zoom 50 --iterations 600 --outputdir result/default/
	mkdir -p result/circle/
	./mandelbrot-cli --zoom 15 --width 300 --height 200 --zoom-config circle:150:100:2:12 --iterations 400 --outputdir result/circle/

