# memo-RPC

## gRPC 环境配置

```shell
$ go get -u google.golang.org/grpc
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

### protoc-gen-go

protoc-gen 是 `.proto` 文件的编译器，可以使用其将 `.proto` 文件生成 go 代码。

## 项目结构

这个 memo 主要分为 用户模块 和 事务模块

```shell
.
├── go.mod
├── go.sum
├── proto
│   ├── pb_go
│   │   └── search.pb.go
│   └── search.proto
├── README.md
└── server.go
```

