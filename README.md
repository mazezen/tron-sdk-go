# tron-sdk-go
[![Go](https://img.shields.io/badge/Go-1.24.1-00ADD8.svg?logo=go&logoColor=white)](https://go.dev/)
[![Go Report Card](https://goreportcard.com/badge/github.com/mazezen/tron-sdk-go)](https://goreportcard.com/report/github.com/mazezen/tron-sdk-go)
[![License: LGPL v3](https://img.shields.io/badge/License-LGPL%20v3-blue.svg)](https://www.gnu.org/licenses/lgpl-3.0)
> tron-sdk-go 是一个的Go SDK用于与TRON区块链进行交互.

## 使用
```shell
go get github.com/mazezen/tron-sdk-go
```

## 安装protoc依赖 protoc-gen-go && protoc-gen-go-grpc
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## protobuf 依赖
* https://github.com/tronprotocol/protocol
* https://github.com/googleapis/googleapis


## 方法
### Block
|   GetNowBlock   |   GetNowBlock2   |   GetBlockByNum   |   GetBlockByNum2   |   GetTransactionCountByBlockNum   |
| ---- | ---- | ---- | ---- | ---- |
|   GetBlockById   |   GetBlockByLimitNext   |   GetBlockByLimitNext2   |   GetBlockByLatestNum   |   GetBlockByLatestNum2   |

