FROM golang as builder

RUN go get -u -v github.com/gocaio/goca
WORKDIR $GOPATH/src/github.com/gocaio/goca/goca

ENV GO111MODULE on
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/goca


FROM scratch

COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/
COPY --from=builder /go/bin/goca /usr/local/bin/goca

ENTRYPOINT ["/usr/local/bin/goca"]