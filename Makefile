NAME ?= bluelytics

all: clean build move start

build: clean
	if [ ! -e ~/.config/infisical/.env ]; then $(MAKE) generate_config; fi
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/$(NAME) src/main.go

generate_config:
	cp .env.example .env
	ln -s $(PWD)/.env ~/.config/infisical/.env

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