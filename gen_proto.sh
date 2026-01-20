#!/bin/bash
set -e

mkdir -p ./pb/tron

COMMON_ARGS=(
  -I=./proto/tron
  -I=./proto/googleapis
  -I=/Users/bz/protobuf/include
)

# 生成所有 TRON proto 到 ./pb/tron/ 目录下
protoc "${COMMON_ARGS[@]}" \
  --go_out=paths=source_relative:./pb/tron \
  --go-grpc_out=paths=source_relative:./pb/tron \
  --go-grpc_opt=require_unimplemented_servers=false \
  ./proto/tron/*.proto


protoc "${COMMON_ARGS[@]}" \
  --go_out=./pb \
  --go_opt=paths=source_relative \
  --go-grpc_out=./pb \
  --go-grpc_opt=paths=source_relative \
  --go-grpc_opt=require_unimplemented_servers=false \
  ./proto/googleapis/google/api/*.proto \
  ./proto/googleapis/google/rpc/*.proto

echo "生成完成..."
echo "TRON 文件在 ./pb/tron/"
echo "Google 文件在 ./pb/google/"