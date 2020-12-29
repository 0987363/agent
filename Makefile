PACKAGE = .

export GO15VENDOREXPERIMENT=1
export GO111MODULE=on

all: build

clean:
	rm -f agent

release: 
	GOARCH=amd64 GOOS=linux go build -v 

build: 
	go build -v

run: build
	./discovery
