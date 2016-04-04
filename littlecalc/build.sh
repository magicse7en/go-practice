#!/usr/bin/env bash

pkg=$1

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

gofmt -w src

go install -x $1

export GOPATH="$OLDGOPATH"

echo 'finished'
