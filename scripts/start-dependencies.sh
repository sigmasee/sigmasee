#!/usr/bin/env sh

set -e
set -x

command=${@:-up -d --build}

cd "$(dirname "${0}")/.."

docker compose -p "sigmasee" \
    --profile core \
    -f docker-compose.yml \
    -f all-in-one/docker-compose.yml \
    --env-file .env \
    $command
