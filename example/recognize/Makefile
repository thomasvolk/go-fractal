
all: build
	
build: test
	go build recognize.go

test:
	(cd ../.. && go test ./...)

clean:
	rm -rf recognize *.png learnset net.json

test-run: build
	mkdir -p learnset
	time ./recognize

