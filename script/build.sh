#!/usr/bin/env bash

set -e
SCRIPT_DIR=$(cd $"${BASH_SOURCE%/*}" && pwd)
BASE_DIR=$(cd $SCRIPT_DIR && cd .. && pwd)
source $SCRIPT_DIR/common.sh

GO_MODULE_PKG='github.com/iwauo/jmx-tools'
COMMAND_NAME='jmx-logger'

set -x

(cd $BASE_DIR \
  && go test "$GO_MODULE_PKG/jmxclient" \
  && go get "$GO_MODULE_PKG/cmd/$COMMAND_NAME"
)