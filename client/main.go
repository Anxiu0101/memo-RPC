package main

import (
	"context"
	"google.golang.org/grpc/credentials"
	"log"
	"memo-RPC/client/pkg/util"
	"time"

	"google.golang.org/grpc"
	pb "memo-RPC/client/ecommerce"
)

const (
	UserPort  = ":9001"
	EventPort = ":9002"
)

var (
	Token string
)

func main() {
	testUserService()
	testEventService()
}

func testUserService() {
	// 创建拨号选项
	var opts []grpc.DialOption
	// 从证书文件中为客户端构造 TLS 凭证
	certs, err := credentials.NewClientTLSFromFile("../certs/memo.pem", "www.anxiu.online")
	if err != nil {
		log.Fatalf("credentials.NewClientTLSFromFile err: %v", err)
	} else {
		opts = append(opts, grpc.WithTransportCredentials(certs))
	}

	// 拨号连接
	//conn, err := grpc.Dial(":"+UserPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(UserPort, opts...)
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	} else {
		log.Println("Success connect User Service")
	}
	defer conn.Close()
	// 创建新客户端，使用拨号创建的连接
	client := pb.NewUserServiceClient(conn)

	// 设置过期时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 访问该函数，获取 响应
	//resp1, err := client.Register(ctx, &pb.UserRegisterRequest{
	//	Username: "Anxiu",
	//	Password: "123456",
	//})
	//if err != nil {
	//	log.Fatalf("client.Register err: %v", err)
	//}
	//log.Printf("Greeting: %s", resp1.String())

	resp2, err := client.Login(ctx, &pb.UserLoginRequest{
		Username: "Anxiu",
		Password: "123456",
	})
	if err != nil {
		log.Fatalf("client.Login err: %v", err)
	}
	log.Printf("Token: %s", resp2.Token)
	Token = resp2.Token
}

func testEventService() {
	// 创建拨号选项
	var opts []grpc.DialOption
	// 从证书文件中为客户端构造 TLS 凭证
	certs, err := credentials.NewClientTLSFromFile("../certs/memo.pem", "www.anxiu.online")
	if err != nil {
		log.Fatalf("credentials.NewClientTLSFromFile err: %v", err)
	} else {
		opts = append(opts, grpc.WithTransportCredentials(certs))
	}

	// 将用户登录获取到的 token 写入 TokenAuth 并传递给服务端
	token := util.TokenAuth{
		Token: Token,
	}
	opts = append(opts, grpc.WithPerRPCCredentials(&token))

	// 拨号连接
	conn, err := grpc.Dial(EventPort, opts...)
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	} else {
		log.Println("Success connect Event Service")
	}
	defer conn.Close()
	// 创建新客户端，使用拨号创建的连接
	client := pb.NewEventServiceClient(conn)

	// 设置过期时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 访问该函数，获取响应
	// 创建新事件
	//log.Println("Before Creating Event")
	//eventResp1, err := client.CreateEvent(ctx, &pb.CreateEventRequest{
	//	Item: &ecommerce.Event{
	//		Name:      "test1",
	//		Content:   "test Event function",
	//		EndTimeAt: time.Now().Add(3 * time.Hour).Unix(),
	//		Type:      1,
	//		State:     1,
	//	},
	//})
	//if err != nil {
	//	log.Fatalf("client.CreateEvent err: %v", err)
	//}
	//log.Printf("Greeting: %s", eventResp1.String())
	//log.Println("After Creating Event")

	resp2, err := client.DeleteEvent(ctx, &pb.DeleteEventRequest{
		Id: "2",
	})
	if err != nil {
		log.Fatalf("client.DeleteEvent err: %v", err)
	}
	log.Printf("Greeting: %s", resp2.String())

	resp3, err := client.UpdateEvent(ctx, &pb.UpdateEventRequest{
		Item: &pb.Event{
			Id:   1,
			Name: "test1 updated",
		},
	})
	if err != nil {
		log.Fatalf("client.UpdateEvent err: %v", err)
	}
	log.Printf("Greeting: %s", resp3.String())

	resp4, err := client.ShowEvent(ctx, &pb.ShowEventRequest{
		Id: "1",
	})
	if err != nil {
		log.Fatalf("client.ShowEvent err: %v", err)
	}
	log.Printf("Greeting: %s", resp4.String())
}
