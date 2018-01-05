#!/bin/sh

rm git-mount
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o git-mount ./main.go
docker build -t a8uhnf/git-mount:1.0.0 .
rm git-mount
docker push a8uhnf/git-mount:1.0.0