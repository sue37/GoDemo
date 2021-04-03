#!/bin/bash

go_main_dir=`pwd`
cpp_lib_dir=$go_main_dir"/../CPyEngine/build"

echo "$go_main_dir"
echo "$cpp_lib_dir"

echo "Start to build"
cd ./CPyEngine/extern/
git clone https://github.com/pybind/pybind11.git
cd pybind11
mkdir build
cd build
cmake ..
cmake --build .
echo "Lib build done"
rm ../libs/libCPyEngine_library.a
ls ../libs/
mv libCPyEngine_library.a ../libs
echo "Lib ready"

echo "Start to build main go"
export SRCDIR=.

cd "$go_main_dir"
rm ./cgo
GOARCH="amd64"
GOOS="linux"
GOPROXY=""
GORACE=""
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GCCGO="gccgo"
CC="gcc"
CXX="g++"
CGO_LDFLAGS+=-lstdc++
GOGCCFLAGS="-fPIC -m64 -pthread -static-libstdc++"
go build cgo.go
