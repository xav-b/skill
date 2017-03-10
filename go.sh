#!/bin/bash

set -e

# const (
  readonly PROJECT=$(basename $(pwd))
  readonly BUILD_DEST="bin"
  readonly TARGETS=("darwin" "linux" "window")
  readonly ARCHS=("386" "amd64")
# )

crosscompile () {
  local dest=${1:-bin}

  test -d ./bin || mkdir bin

  for GOOS in ${TARGETS}; do
    for GOARCH in ${ARCHS}; do
      echo "Building $GOOS-$GOARCH"
      export GOOS=$GOOS
      export GOARCH=$GOARCH

      go build -o ${dest}/${PROJECT}-$GOOS-$GOARCH
    done
  done
}

if [ "$#" -lt 1 ]
then
    echo "No command provided."
    exit 1
fi

case "$1" in
  cross) echo "Cross compiling..."
    # TODO pass dest, os and arch as args
    crosscompile "$(BUILD_DEST)"
    ;;
  version)
    go version
    ;;
  *) echo "Invalid command, see https://github.com/hackliff/skill for reference."
    ;;
esac
exit 0
