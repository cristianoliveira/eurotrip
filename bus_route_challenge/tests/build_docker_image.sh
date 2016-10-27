#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
pushd $DIR > /dev/null
if [[ $1 == *.zip ]]; then
    echo "zip mode"
    mkdir -p docker/files
    cp $1 docker/files/
    T="zip"
    F=`ls docker/files/`
else
    echo "git mode"
    T="git"
    F=$1
fi
docker build -t dev-test --build-arg path=$F --build-arg type=$T  docker
rm -fr docker/files
popd > /dev/null
