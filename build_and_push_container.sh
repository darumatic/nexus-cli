#!/usr/bin/env bash
env GOOS=linux GOARCH=amd64 go build &&
docker build . -t darumatic/nexus-cleaner:3.0 &&
docker push darumatic/nexus-cleaner:3.0