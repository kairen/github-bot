# Build stage
FROM golang:1.8-alpine AS build-env

ENV GOPATH "/go"

ADD . /go/src/bot
RUN apk add --no-cache git && \
  go get -u github.com/golang/dep/cmd/dep
RUN cd /go/src/bot && \
  dep ensure && \
  go build -o bot

# Run stage
FROM alpine
MAINTAINER Kyle Bai(kyle.b@inwinstack.com)

COPY --from=build-env /go/src/bot/bot /bin/bot

ENTRYPOINT ["/bin/bot"]
