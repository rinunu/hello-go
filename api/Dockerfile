# create image from the official Go image
FROM golang:1.10.0-alpine3.7

RUN apk add --update tzdata \
    bash wget curl git;

RUN go get github.com/pilu/fresh
RUN go get github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/rinunu/hello-go/api
VOLUME /go/src/github.com/rinunu/hello-go/api

ENV GOPATH /go  

# serve the app
CMD fresh -c runner.conf main.go
