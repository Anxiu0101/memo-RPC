## 依赖

```shell
# gRPC 依赖
go get -u google.golang.org/grpc
# gRPC 中 .proto 文件的编译插件
go get -u github.com/golang/protobuf/protoc-gen-go
# gRPC gateway 用于转化 RESTful 请求为 protobuf 格式
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
```

```shell
├─client
│  ├─ecommerce
│  └─router
└─server
    ├─certs
    ├─conf
    ├─ecommerce
    ├─model
    ├─pkg
    │  └─util
    ├─proto
    │  └─google
    │      └─api
    └─service
```



### 模型

![image-20220529231530039](https://raw.githubusercontent.com/Anxiu0101/PicgoImg/master/202205292315060.png)

![image-20220529231530039](https://raw.githubusercontent.com/Anxiu0101/PicgoImg/master/202205310720330.png)



## 实现思路

通过 gRPC 的 gateway 进行 RESTful API的暴露，将此部分作为项目的前台系统，前台系统通过 RPC 对后台服务进行调用，达成分布式服务的要求。
所以现在就是将原本的项目拆分为前台与后台项目，前台项目将来自网络的请求转化为 proto 请求发送至后台，是原项目的 API 层，后台便是原项目中的 service 层，与数据库进行交互。
那么如何对用户服务和任务服务进行拆分呢？

## 关于 Protoc

```shell
protoc -I ../../ -I ./ --go_out=. *.proto
```

Protoc 即编译 `.proto` 文件的工具， 其中有几个重要参数

```shell
$ protoc --help
  -IPATH, --proto_path=PATH   Specify the directory in which to search for
                              imports.  May be specified multiple times;
                              directories will be searched in order.  If not
                              given, the current working directory is used.
                              If not found in any of the these directories,
                              the --descriptor_set_in descriptors will be
                              checked for required proto file.
  --plugin=EXECUTABLE         Specifies a plugin executable to use.
                              Normally, protoc searches the PATH for
                              plugins, but you may specify additional
                              executables not in the path using this flag.
                              Additionally, EXECUTABLE may be of the form
                              NAME=PATH, in which case the given plugin name
                              is mapped to the given executable even if
                              the executable's own name differs.
  --go_out=OUT_DIR            Generate Go header and source.

```

需要注意的是，`.proto` 文件中的 ==import== 关键字是基于项目根目录的，因此在使用 protoc 指令时也需要将 `-IPATH` 调整到根目录。

## 关于 OpenSSL

```shell
$ openssl genrsa -out server/server-key.pem
$ openssl req -new -out server/server-req.csr -key server/server-key.pem

You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [AU]:CN
State or Province Name (full name) [Some-State]:fujian
Locality Name (eg, city) []:fuzhou
Organization Name (eg, company) [Internet Widgits Pty Ltd]:Anxiu
Organizational Unit Name (eg, section) []:Anxiu
Common Name (e.g. server FQDN or YOUR name) []:memo-RPC
Email Address []:anxiu.fyc@foxmail.com

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:Ax7391839@ssl
An optional company name []:Anxiu
```



## 关于 Go-callsiv

```shell
go-callvis -debug -tests E:\Desktop\West2Go\6\memo-RPC\server\ecommerce\
# 将在 localhost:7878 展示
```



关于 微服务 我现在的理解是，将原本一个大型的项目进行拆分，分解成几个微服务。在服务器集群中，每一台主机只负责其中一个微服务模块。而 RPC 是解决不同服务之间的函数调用的渠道。

## 关于用户认证

在这个 demo 中使用 gRPC 的拦截器与 token 实现对于用户操作权限的鉴权。

