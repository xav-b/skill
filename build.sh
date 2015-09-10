#!/bin/sh

GIT_USER="$(git config --get user.name)"
PROJECT="$( basename $(pwd) )"
SOURCES_PATH="/go/src/github.com/${GIT_USER}/${PROJECT}"

# NOTE Does the --rm remove volumes ?
# TODO Check $@ and default to build
docker run -it --rm --volume $PWD:$SOURCES_PATH --workdir ${SOURCES_PATH} golang go $@
echo "result: $?"
#sudo chown -R $USER .
