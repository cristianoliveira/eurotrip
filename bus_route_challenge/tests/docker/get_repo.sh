#!/bin/bash

if [ $1 == "git" ]; then
    git clone $2
elif [ $1 == "zip" ]; then
    unzip -d . /files/$2
fi
