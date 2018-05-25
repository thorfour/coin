.PHONY: all clean gcp

EXECUTABLE='coin'
ZIP='coin.zip'
PLUG="coin"

all: 
	go build -o $(EXECUTABLE) ./cmd/serverless
clean:
	rm $(EXECUTABLE)
	rm $(ZIP)
gcp: all
	zip -j $(ZIP) ./build/gcp/* $(EXECUTABLE)
plugin: 
	mkdir -p ./bin/plugin
	CGO_ENABLED=1 GOOS=linux go build -buildmode=plugin -o ./bin/plugin/$(PLUG) ./cmd/plugin/
