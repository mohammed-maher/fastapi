from golang:latest as builder

COPY ./cmd/fastapi/ /app

WORKDIR /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -command="./app" -build="go build -o app ."



