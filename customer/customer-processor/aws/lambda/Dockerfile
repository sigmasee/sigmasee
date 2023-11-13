# syntax = docker/dockerfile:1.4.1
ARG FINAL_IMAGE=alpine:3
ARG BUILD_IMAGE=golang:1.21

FROM $BUILD_IMAGE AS build

RUN mkdir -p /src/shared/enterprise \
    mkdir -p /src/shared/clients \
    mkdir -p /src/customer/shared \
    mkdir -p /src/customer/customer-processor

COPY go.mod /src/
COPY go.sum /src/
COPY shared/enterprise /src/shared/enterprise
COPY shared/clients /src/shared/clients
COPY customer/shared /src/customer/shared
COPY customer/customer-processor /src/customer/customer-processor

WORKDIR /src/customer/customer-processor/aws/lambda/customer-v1-event
RUN --mount=type=cache,target=/go GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bootstrap -tags lambda.norpc -ldflags "-s -w" main.go

FROM build AS test

WORKDIR /src/customer/customer-processor/aws/lambda

RUN mkdir -p /coverage/reports
RUN --mount=type=cache,target=/go GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go test -coverpkg=./... -v -covermode=count -coverprofile="/coverage/reports/coverage.out" ./...

FROM $FINAL_IMAGE AS final
LABEL maintainer="morteza.alizadeh@gmail.com"

RUN apk update && \
    apk add --upgrade apk-tools && \
    apk upgrade --available

WORKDIR /sigmasee/customer-v1-event
COPY --from=build /src/customer/customer-processor/aws/lambda/customer-v1-event/bootstrap /sigmasee/customer-v1-event
COPY --from=build /src/customer/customer-processor/aws/lambda/customer-v1-event/config.yaml /sigmasee/customer-v1-event
