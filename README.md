# go-fractal

Fractal library for go lang.

## Examples

build example app:

    cd example
    make

generate png:

    ./mandelbrot-cli -xstart 0.435396403 -xend 0.451687191 -ystart 0.367981352 -yend 0.380210061 -width 1200 -height 400 -iterations 100

start server:

    ./mandelbrot-cli -serve

* http://localhost:8080/
* http://localhost:8080/?xstart=0.435396403&xend=0.451687191&ystart=0.367981352&yend=0.380210061&width=200
* http://localhost:8080/?xstart=-0.7515104166666667&xend=-0.7384895833333335&ystart=-0.1179174885797342&yend=-0.1081518635797342


