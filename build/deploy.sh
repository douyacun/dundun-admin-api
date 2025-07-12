#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && cd .. && pwd )"
echo $DIR
NAME=ad-materials

function build() {
  echo "GOOS=linux GOARCH=amd64 go build -ldflags -s -w -o $DIR/build/bin/$NAME $DIR/main.go"
  GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $DIR/build/bin/$NAME $DIR/main.go
  if [ $? -ne 0 ]; then
    echo "build binary failed"
    exit -1
  fi
}

function clear() {
    rm -rf $DIR/build/bin
}

clear && build 
