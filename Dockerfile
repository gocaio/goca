FROM golang:alpine

RUN apk add gcc git musl-dev && go get -u -v github.com/gocaio/goca

WORKDIR /go/src/github.com/gocaio/goca

ENV GO111MODULE on
RUN go get ./...

ENTRYPOINT ["go","run","goca/goca.go"]
