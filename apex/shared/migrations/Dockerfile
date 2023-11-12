# syntax = docker/dockerfile:1.4.1
ARG BUILD_IMAGE=ubuntu:latest

FROM $BUILD_IMAGE AS test

RUN mkdir -p /coverage/reports/

FROM $BUILD_IMAGE AS final
LABEL maintainer="morteza.alizadeh@gmail.com"

RUN apt-get update -y \
  && apt-get upgrade -y \
  && apt-get install curl -y \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /
RUN curl https://release.ariga.io/atlas/atlas-linux-amd64-latest -o /atlas
RUN chmod +x /atlas

WORKDIR /migrations
COPY apex/shared/migrations /migrations

CMD /atlas migrate apply --dir file:///migrations/ --url $(echo "$sigmasee_POSTGRES_CONNECTIONSTRING" | tr -d '"')
