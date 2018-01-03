.PHONY: all clean gcp

EXECUTABLE='coin'
ZIP='coin.zip'

all: 
	go build -o $(EXECUTABLE) ./cmd/
clean:
	rm $(EXECUTABLE)
	rm $(ZIP)
gcp: all
	zip -j $(ZIP) ./build/gcp/* $(EXECUTABLE)
