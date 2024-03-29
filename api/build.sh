#!/usr/bin/env sh

if [ -z "$E_DEBUG" ]; then
    GOOS=linux go build -o /api github.com/tetrago/motmot/api/cmd/api
else
    CGO_ENABLE=0 GOOS=linux go build -tags debug -o /api github.com/tetrago/motmot/api/cmd/api
fi