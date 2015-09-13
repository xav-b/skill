#!/bin/bash

set -e

# const (
  readonly PROJECT=$(basename $(pwd))
# )

if [ "$#" -lt 1 ]
then
    echo "No command provided."
    exit 1
fi

crosscompile () {
  local dest=${1:-bin}
  test -d ./bin || mkdir bin
  for GOOS in darwin linux windows; do
    for GOARCH in 386 amd64; do
      echo "Building $GOOS-$GOARCH"
      export GOOS=$GOOS
      export GOARCH=$GOARCH
      go build -o ${dest}/${PROJECT}-$GOOS-$GOARCH
    done
  done
}

case "$1" in
  cross) echo "Cross compiling..."
    # TODO pass dest, os and arch as args
    crosscompile "bin"
    ;;
  version)
    go version
    ;;
  *) echo "Invalid command, see https://github.com/hackliff/skill for reference."
    ;;
esac
exit 0
