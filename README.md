## 依赖

```shell
# gRPC 依赖
go get -u google.golang.org/grpc
# gRPC 中 .proto 文件的编译插件
go get -u github.com/golang/protobuf/protoc-gen-go
# gRPC gateway 用于转化 RESTful 请求为 protobuf 格式
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
```



### 模型

![image-20220529231530039](https://raw.githubusercontent.com/Anxiu0101/PicgoImg/master/202205292315060.png)



## 实现思路

通过 gRPC 的 gateway 进行 RESTful API的暴露，将此部分作为项目的前台系统，前台系统通过 RPC 对后台服务进行调用，达成分布式服务的要求。
所以现在就是将原本的项目拆分为前台与后台项目，前台项目将来自网络的请求转化为 proto 请求发送至后台，是原项目的 API 层，后台便是原项目中的 service 层，与数据库进行交互。
那么如何对用户服务和任务服务进行拆分呢？