SOURCES := $(shell find . -name '*.go')
BINARY 	:= API-go

all: build

build: $(SOURCES)
	go build

run: all
	./API-go

clean:
	rm -f $(BINARY)

.PHONY: run clean docker build-docker-image

test:
	go test ./...
