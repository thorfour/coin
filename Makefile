.PHONY: all clean gcp plugin server docker

GO_VERSION="1.11"
EXECUTABLE='coin'
ZIP='coin.zip'
PLUG="coin"

go = @docker run \
        --rm \
        -v ${PWD}:/go/src/github.com/thorfour/coin \
        -w /go/src/github.com/thorfour/coin \
        -u $$(id -u) \
        -e XDG_CACHE_HOME=/tmp/.cache \
        -e CGO_ENABLED=0 \
        -e GOOS=linux \
        -e GOPATH=/go \
        golang:$(GO_VERSION) \
        go

go_cgo = @docker run \
        --rm \
        -v ${PWD}:${PWD} \
        -w ${PWD} \
        -u $$(id -u) \
        -e XDG_CACHE_HOME=/tmp/.cache \
        -e CGO_ENABLED=1 \
        -e GOOS=linux \
        -e GOPATH=/go \
        golang:$(GO_VERSION) \
        go

all: serverless docker plugin

serverless:
	mkdir -p ./bin/serverless
	$(go) build -o ./bin/serverless/$(EXECUTABLE) ./cmd/serverless

gcp: all
	mkdir -p ./bin/gcp
	zip -j $(ZIP) ./build/gcp/* ./bin/gcp/$(EXECUTABLE)

plugin: 
	mkdir -p ./bin/plugin
	$(go_cgo) build -buildmode=plugin -o ./bin/plugin/$(PLUG) ./cmd/plugin/

docker: 
	docker build -t quay.io/thorfour/coin .

server:
	mkdir -p ./bin/server
	$(go) build -o ./bin/server/$(EXECUTABLE) ./cmd/server

clean:
	rm -rf ./bin	
