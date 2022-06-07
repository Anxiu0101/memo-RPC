[TOC]

## 1. 介绍

将原本的备忘录项目进行拆分，分为三个部分

第一个是备忘录系统的 client，也是系统的前台。主要功能是调用另外两个微服务模块的服务，另外也通过 `gRPC-gateway` 将 `RESTful` 请求转化为 `proto` 请求，使得该系统可以同时处理 RESTful 协议和 RPC 协议的请求。

第二个是备忘录系统的用户服务，主要针对的是用户相关的请求。这一块的主要技术要点是 gRPC 的 token 颁发。端口为 9002

第三个是备忘录系统的事件服务，也是备忘录系统的主要服务。这一块的主要技术要点是通过 gRPC 拦截器进行 token 的认证，以及具体的业务代码。端口为 9001


## 2. 任务

- [ ] 完成客户端和服务端的 TLS 连接设置
- [ ] 完成客户端的 token 添加 和 服务端的拦截器验证 token 设置
- [ ] 完成备忘录相关的业务代码


## 3. 思考

我认为我这次的结构是有问题的，RPC 本身是用来处理微服务之间的调用的，微服务解耦拆开之后并不是完全不相关的，服务之间还存在着调用关系，所以按理来说，应该三个微服务项目都保存了 `.pb.go` 的存根，以满足用户服务与备忘录服务之间互相调用的函数。

### 3.1. 关于事件的持有者

每一个备忘录的事件都为创建其的用户所拥有，这是一个一对多的关系，


## 4. 依赖

```shell
# gRPC 依赖
go get -u google.golang.org/grpc
# gRPC 中 .proto 文件的编译插件
go get -u github.com/golang/protobuf/protoc-gen-go
# gRPC gateway 用于转化 RESTful 请求为 protobuf 格式
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
```

```shell
├───certs       # TLS .key .csr .pem file 
├───client      # 前台服务系统
│   ├───ecommerce # .pb.go file
│   ├───router 
│   └───main.go
├───eventserver # 备忘录事件服务系统
│   ├───conf        # 配置文件，使用 viper
│   ├───ecommerce   # pb.go 存根
│   ├───model       # 模型
│   ├───pkg         # 第三方包
│   │   └───util
│   ├───proto       # IDF 接口定义文件
│   │   └───google
│   │       └───api # 用于支持 gateway 的两个 .proto 文件
│   ├───service     # 业务逻辑代码
│   └───main.go
└───userserver  # 用户服务系统
    └───...

```









### 4.1. 模型

![image-20220529231530039](https://raw.githubusercontent.com/Anxiu0101/PicgoImg/master/202205292315060.png)

![image-20220529231530039](https://raw.githubusercontent.com/Anxiu0101/PicgoImg/master/202205310720330.png)



## 5. 实现思路

通过 gRPC 的 gateway 进行 RESTful API的暴露，将此部分作为项目的前台系统，前台系统通过 RPC 对后台服务进行调用，达成分布式服务的要求。
所以现在就是将原本的项目拆分为前台与后台项目，前台项目将来自网络的请求转化为 proto 请求发送至后台，是原项目的 API 层，后台便是原项目中的 service 层，与数据库进行交互。
那么如何对用户服务和任务服务进行拆分呢？

## 6. 关于 Protoc

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

## 7. 关于 OpenSSL

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



## 8. 关于 Go-callsiv

```shell
go-callvis -debug -tests E:\Desktop\West2Go\6\memo-RPC\server\ecommerce\
# 将在 localhost:7878 展示
```


## 9. 关于用户认证

在这个 demo 中使用 gRPC 的拦截器与 token 实现对于用户操作权限的鉴权。

在服务端入口增加拦截器代码

```go
	// 使用一元拦截器（grpc.UnaryInterceptor），验证请求
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 拦截普通方法请求，验证Token
		err = service.CheckAuthority(ctx)
		if err != nil {
			return
		}
		// 继续处理请求
		return handler(ctx, req)
	}
```

