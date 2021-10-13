FROM golang:1.17.2-alpine3.14 AS builder
WORKDIR /go/src
COPY go.mod .
COPY main.go .
RUN go mod vendor
ENV CGO_ENABLED=0
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/httpserver main.go

FROM alpine:3.14
LABEL name="httpserver"
WORKDIR /opt/httpserver
COPY --form=builder /go/bin/httpserver /opt/httpserver/
EXPOSE 8080
ENTRYPOINT  ["/opt/httpserver/httpserver"]
