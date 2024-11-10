NAME ?= bluelytics

all: clean build move start

dev:
	go run src/main.go

build: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/$(NAME) src/main.go

start:
	$(NAME)

clean:
	rm -f bin/$(NAME)

fmt:
	go fmt ./...

test:
	go test -v ./...

move:
	sudo mv bin/$(NAME) /usr/local/bin

.PHONY: clean