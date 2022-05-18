# memo-RPC

## gRPC 环境配置

```shell
$ go get -u google.golang.org/grpcgo get -u google.golang.org/grpc
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

### protoc-gen-go

protoc-gen 是 `.proto` 文件的编译器，可以使用其将 `.proto` 文件生成 go 代码。
