#!/usr/bin/env sh

set -e
set -x

cd /src/shared/sigmaseectl

until [ ]; do
    ./sigmaseectl database provision --name customer && break
    sleep 1
done

./sigmaseectl database provision --name apex

/atlas migrate apply --dir file:///customer/migrations/ --url postgresql://root@cockroachdb.localhost:26257/customer?sslmode=disable
/atlas migrate apply --dir file:///apex/migrations/ --url postgresql://root@cockroachdb.localhost:26257/apex?sslmode=disable
