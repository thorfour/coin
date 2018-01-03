.PHONY: all

EXECUTABLE='coin'

all: 
	go build -o $(EXECUTABLE) ./cmd/
clean:
	rm $(EXECUTABLE)
