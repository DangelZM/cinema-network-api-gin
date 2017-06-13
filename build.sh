#!/bin/bash

which docker && docker --version | grep "Docker version"

if [ "$?" == "0" ]; then
        echo "Docker installed +"
        echo "Building..."

        docker run --rm -v "$PWD":/go/src/github.com/dangelzm/cinema-network-api -w /go/src/github.com/dangelzm/cinema-network-api iron/go:dev go get github.com/tools/godep && godep go build -o app

        echo "Done"
else
    echo "You need Docker install" >&2
fi

