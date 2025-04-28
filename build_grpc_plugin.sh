#!/bin/bash

# This script builds the gRPC PHP plugin and installs the Go protobuf compiler plugin.
# Both are required for protoc to generate gRPC code.

SCRIPTPATH="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )";
tmpDir="$SCRIPTPATH/tmp";
protocDir="$SCRIPTPATH/protoc";

if [ -d "$tmpDir" ]; then
    rm -rf "$tmpDir"
fi


mkdir "$tmpDir" && \
cd "$tmpDir" && \
git clone https://github.com/grpc/grpc "$tmpDir" && \
cd grpc && \
git submodule update --init && \
mkdir -p cmake/build && \
cd cmake/build && \
cmake ../.. && \
make protoc grpc_php_plugin && \
cp "$tmpDir/cmake/build/grpc_php_plugin" "$protocDir" && \
cd "$tmpDir" && git clone https://github.com/grpc/grpc-go && \
cd grpc-go && \
git checkout v1.57.2 && \
cd cmd/protoc-gen-go-grpc && \
go build -o "$protocDir/protoc_gen_go_grpc" && \
rm -rf "$tmpDir" && \
echo -e "\n******************\n\tDONE\n******************\n"