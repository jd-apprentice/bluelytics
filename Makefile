NAME ?= bluelytics

all: clean build move start

build: clean
	go build -o bin/$(NAME) main.go

start:
	$(NAME)

clean:
	rm -f bin/$(NAME)

move:
	sudo mv bin/$(NAME) /usr/local/bin

.PHONY: clean