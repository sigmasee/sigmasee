# syntax = docker/dockerfile:1.4.1
ARG BUILD_IMAGE=golang:1.21

FROM $BUILD_IMAGE AS build
LABEL maintainer="morteza.alizadeh@gmail.com"

RUN apt-get update -y
RUN apt-get install unzip

ENV GIT_TAG="25.0"
ENV PROTOC_ZIP=protoc-$GIT_TAG-linux-x86_64.zip
RUN curl -OL https://github.com/google/protobuf/releases/download/v$GIT_TAG/$PROTOC_ZIP
RUN unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
RUN unzip -o $PROTOC_ZIP -d /usr/local include/*
RUN rm -f $PROTOC_ZIP

RUN go install github.com/golang/protobuf/protoc-gen-go@latest

RUN apt-get clean

COPY go.mod /
COPY go.sum /
COPY shared/enterprise /shared/enterprise
COPY shared/clients /shared/clients
COPY shared/sigmaseectl /shared/sigmaseectl

WORKDIR /shared/sigmaseectl
RUN --mount=type=cache,target=/go CGO_ENABLED=0 go build -v -o /bin/sigmaseectl main.go

RUN mkdir -p /output
COPY api-definitions/events /events

RUN protoc \
    -I=/events/sigmasee \
    --go_out=/output \
    customer_v1.proto

RUN /bin/sigmaseectl \
    generate \
    client \
    event \
    --protobufFilePath=/events/sigmasee/customer_v1.proto \
    --packageName=v1 \
    --eventType=Event \
    --topicName=customer.v1.event \
    --retryTopicNamePrefix=customer.v1.event.retry \
    --retryTopicNameCount 1 \
    --deadLetterTopicName=customer.v1.event.deadletter \
    --outputPath=/output/shared/clients/events/sigmasee/customer/v1
