# syntax = docker/dockerfile:1.4.1
ARG FINAL_IMAGE=alpine:3
ARG BUILD_IMAGE=golang:1.21

FROM $BUILD_IMAGE AS build

RUN mkdir -p /src/shared/enterprise \
    mkdir -p /src/shared/clients \
    mkdir -p /src/apex/shared \
    mkdir -p /src/apex/apex-api

COPY go.mod /src/
COPY go.sum /src/
COPY shared/enterprise /src/shared/enterprise
COPY shared/clients /src/shared/clients
COPY apex/shared /src/apex/shared
COPY apex/apex-api /src/apex/apex-api

WORKDIR /src/apex/apex-api

RUN --mount=type=cache,target=/go CGO_ENABLED=0 go build -v -o /bin/apex-api main.go

FROM build AS test

RUN mkdir -p /coverage/reports
RUN --mount=type=cache,target=/go CGO_ENABLED=0 go test -coverpkg=./... -v -covermode=count -coverprofile="/coverage/reports/coverage.out" ./...

FROM $FINAL_IMAGE AS final
LABEL maintainer="morteza.alizadeh@gmail.com"

RUN apk update && \
    apk add --upgrade apk-tools && \
    apk upgrade --available
RUN apk --no-cache add curl

WORKDIR /sigmasee
COPY --from=build /bin/apex-api /sigmasee
COPY --from=build /src/apex/apex-api/config.yaml /sigmasee

CMD ["/sigmasee/apex-api", "start"]
