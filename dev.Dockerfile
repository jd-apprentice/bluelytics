FROM golang:1.23 AS development

ARG ENVIRONMENT=DEV

WORKDIR /app

RUN apt-get update && apt-get install -y ca-certificates --no-install-recommends
COPY go.mod go.sum ./
RUN go mod download
COPY ./src ./src

ENTRYPOINT ["go", "run", "src/main.go"]