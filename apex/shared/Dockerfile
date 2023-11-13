# syntax = docker/dockerfile:1.4.1
ARG BUILD_IMAGE=golang:1.21

FROM $BUILD_IMAGE AS test
LABEL maintainer="morteza.alizadeh@gmail.com"

RUN mkdir -p /src/shared/enterprise \
    mkdir -p /src/apex/shared

COPY go.mod /src/
COPY go.sum /src/
COPY shared/enterprise /src/shared/enterprise
COPY shared/clients /src/shared/clients
COPY apex/shared /src/apex/shared

WORKDIR /src/apex/shared

RUN mkdir -p /coverage/reports
RUN --mount=type=cache,target=/go CGO_ENABLED=0 go test -coverpkg=./... -v -covermode=count -coverprofile="/coverage/reports/coverage.out" ./...
