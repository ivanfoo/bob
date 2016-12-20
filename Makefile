SHELL = /bin/bash

dependencies:
	go get -v -t ./...

build:
	go build -o build/bob

clean:
	rm -rf build/

