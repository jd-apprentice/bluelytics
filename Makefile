NAME ?= bluelytics

build: clean
	go build -o bin/$(NAME) main.go

start:
	./bin/$(NAME)

clean:
	rm -f bin/$(NAME)

.PHONY: clean