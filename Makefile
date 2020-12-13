all:	mac-app

mac-app:	build
	mkdir -p PhotoSquared.app/Contents/MacOS
	cp $(GOPATH)/bin/photosquared PhotoSquared.app/Contents/MacOS

build:	
	go build -o $(GOPATH)/bin/photosquared main.go
