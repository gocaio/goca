ARG GOOS=linux
ARG GOARCH=amd64

FROM golang:latest AS builder
ADD . /goca/
WORKDIR /goca/
ENV GOOS=${GOOS}
ENV GOARCH=${GOARCH}
RUN go mod download
RUN CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -a -o /goca .

# final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /goca ./
RUN chmod +x ./goca
ENTRYPOINT ["./goca"]
