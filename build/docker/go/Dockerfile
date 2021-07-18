FROM golang:latest

WORKDIR /go/authBeatsPro

COPY ./ /go/authBeatsPro

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build cmd/app/main.go" --command=./main