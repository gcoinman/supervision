#!/bin/sh

SCRIPT_DIR="$(cd "$(dirname "${0}")" && echo "${PWD}")"
${SCRIPT_DIR}/clean.sh

# Go
if  type go > /dev/null 2>&1; then
    protoc --go_out=plugins=grpc:. *.proto
    echo "Create Go file from .proto"
else
    echo "Info: go not found. See README.md"
fi