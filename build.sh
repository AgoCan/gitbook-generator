#!/bin/bash

BASH_NAME=$0
SCRIPT_DIR=$(cd $(dirname ${BASH_SOURCE[0]}); pwd)

docker build -t hank997/gitbook-generator:0.0.1 .
docker run -it --rm -v $SCRIPT_DIR:$SCRIPT_DIR hank997/gitbook-generator:0.0.1 -a hank -k hankbook -n title -p $SCRIPT_DIR
