all:	build

build:	
	go build -o $(GOPATH)/bin/photosquared main.go
