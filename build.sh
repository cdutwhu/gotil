#!/bin/bash
set -e

GOARCH=amd64
LDFLAGS="-s -w"
OUT=main
CGO_ENABLED=0 GOOS="windows" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o $OUT.exe # linux windows darwin
echo "built"