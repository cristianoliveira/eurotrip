#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
pushd $DIR > /dev/null
docker run -d --name=dev-test -p 8088:8088 dev-test > /dev/null 
sleep 10
bash simple_test.sh
docker rm -f dev-test > /dev/null 
popd > /dev/null
