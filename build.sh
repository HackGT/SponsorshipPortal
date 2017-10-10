#!/usr/bin/env bash
set -euxo pipefail

# install and build webpack
(
    cd frontend/
    npm install -y
    npm run build
)

# Move around weird go paths
APP_PATH="${GOPATH}"/src/github.com/HackGT/SponsorshipPortal/
mkdir -p "$APP_PATH"
cp -R backend "$APP_PATH"
(
    cd "$APP_PATH"
    go get ./...
)

revel package github.com/HackGT/SponsorshipPortal/backend prod

