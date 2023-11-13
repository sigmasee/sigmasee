# syntax = docker/dockerfile:1.4.1
ARG BUILD_IMAGE=golang:1.21

FROM $BUILD_IMAGE AS build

RUN mkdir -p /api-definitions \
    mkdir -p /src/shared/enterprise \
    mkdir -p /src/shared/clients \
    mkdir -p /src/all-in-one \
    mkdir -p /src/gateway \
    mkdir -p /src/customer \
    mkdir -p /src/apex

COPY api-definitions /src/api-definitions
COPY go.mod /src/
COPY go.sum /src/
COPY shared/enterprise /src/shared/enterprise
COPY shared/clients /src/shared/clients
COPY all-in-one /src/all-in-one
COPY gateway /src/gateway
COPY customer /src/customer
COPY apex /src/apex

WORKDIR /src/all-in-one

CMD ["go", "run", "main.go", "start"]
