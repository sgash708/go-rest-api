FROM golang:1.16.1-alpine as builder

WORKDIR /docker
COPY ../../  ./
RUN go mod download

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -ldflags '-s -w' ./cmd/go-rest-api

FROM alpine
COPY --from=builder /docker/go-rest-api /opt/app/
ENTRYPOINT ["/opt/app/go-rest-api"]