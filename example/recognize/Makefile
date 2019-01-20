
all: build
	
build: test
	go build recognize.go

test:
	(cd ../.. && go test ./...)

clean:
	rm -rf recognize *.png learnset net.json zoom

test-run: build
	time ./recognize

