
all: build
	
build: test
	go build recognize.go

test:
	(cd ../.. && go test)

clean:
	rm -rf recognize *.png learnset

test-run: build
	mkdir -p learnset
	./recognize

