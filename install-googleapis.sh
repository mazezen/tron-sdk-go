#!/bin/zsh

if [ ! -d "$(pwd)/proto/googleapis/google/api" ]; then
  mkdir -p $(pwd)/proto/googleapis/google/api
fi

if [ ! -d "$(pwd)/proto/googleapis/google/rpc" ]; then
  mkdir -p $(pwd)/proto/googleapis/google/rpc
fi

wget -O $(pwd)/proto/googleapis/google/api/http.proto https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto
wget -O $(pwd)/proto/googleapis/google/api/annotations.proto https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto
wget -O $(pwd)/proto/googleapis/google/rpc/code.proto https://raw.githubusercontent.com/googleapis/googleapis/master/google/rpc/code.proto
wget -O $(pwd)/proto/googleapis/google/rpc/error_details.proto https://raw.githubusercontent.com/googleapis/googleapis/master/google/rpc/error_details.proto
wget -O $(pwd)/proto/googleapis/google/rpc/status.proto https://raw.githubusercontent.com/googleapis/googleapis/master/google/rpc/status.proto


ls proto/googleapis/google
