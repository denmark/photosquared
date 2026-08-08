all:	build

build:
	go build -o bin/photosquared main.go

install:
	go install .
