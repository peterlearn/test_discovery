# Go parameters
GOCMD=GO111MODULE=on CGO_ENABLED=0 go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

all: test build
build:
	rm -rf dist/
	mkdir -p dist/conf
	$(GOBUILD) -o dist/bin/discovery main.go discovery.go

test:
	$(GOTEST) -v ./...

clean:
	rm -rf dist/

run:
	nohup dist/bin/discovery -conf dist/conf -confkey discovery.toml -log.dir dist/log & > dist/nohup.out

stop:
	pkill -f dist/bin/discovery
