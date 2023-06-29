#!/bin/sh
go build -ldflags="-X main.version=0.1.0 -X main.commit=$(git rev-parse --short HEAD) -X main.date=$(date +%Y-%m-%dT%T%z)"
