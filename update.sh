#!/bin/bash

WORKSPACE=$(pwd)
SRC=${WORKSPACE}/raylib

mkdir build
cd build

# Checkout original sources
# TODO: Allow to pull a specific branch
git clone https://github.com/raysan5/raylib.git
cd raylib/src
git pull

# copy files
for f in $(find . -type f); do
				dest="$SRC/$(echo $f | cut -c 3-)"
				echo "Copy $f -> $dest"
				cp -f $f $dest
done

cd $WORKSPACE

# Try to build updated project
go get golang.org/x/tools/cmd/goimports
chmod +x build.sh
./build.sh
