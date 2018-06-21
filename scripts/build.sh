#!/bin/sh

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

go build -ldflags '-extldflags "-static"' -a -v -o ../bin/linux/lark ./cmd/