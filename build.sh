#!/bin/bash
rm -rf libs2go.*
export GOARCH="amd64"
export GOBIN="/usr/local/go/bin"
export GOEXE=""
export GOHOSTARCH="amd64"
export GOHOSTOS="linux"
export GOOS="linux"
export GOPATH="/home/wayhood/Works/golib:/home/wayhood/Works/go"
export GORACE=""
export GOROOT="/usr/local/go"
export GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
export GCCGO="gccgo"
export GO386=""
export CC="gcc"
export GOGCCFLAGS="-fPIC -m32 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build128906296=/tmp/go-build -gno-record-gcc-switches"
export CXX="g++"
export CGO_ENABLED="1"
export PKG_CONFIG="pkg-config"
export CGO_CFLAGS="-g -O2"
export CGO_CPPFLAGS=""
export CGO_CXXFLAGS="-g -O2"
export CGO_FFLAGS="-g -O2"
export CGO_LDFLAGS="-g -O2"

$GOBIN/go build -x -v -ldflags "-s -w" -buildmode=c-shared -o libs2go.so s2go.go

sudo cp libs2go.so /usr/local/lib
sudo cp libs2go.h /usr/local/include
