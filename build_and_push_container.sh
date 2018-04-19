#!/usr/bin/env bash
env GOOS=linux GOARCH=amd64 go build &&
docker build . -t darumatic/nexus-cleaner:latest &&
docker push darumatic/nexus-cleaner:latest