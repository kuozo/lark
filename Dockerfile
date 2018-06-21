FROM golang:1.10.3-alpine as builder
WORKDIR /go/src/github.com/klnchu/lark
COPY . .
RUN ./scripts/build.sh

FROM alpine:latest
LABEL MAINTAINER="Kollin <kollinchu@gmail.com>"

RUN apk --no-cache add ca-certificates
COPY --from=0 /go/src/github.com/klnchu/lark/bin/linux/lark /usr/local/bin/lark