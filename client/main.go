package main

import (
	"context"
	"log"
	"memo-RPC/client/ecommerce"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "memo-RPC/client/ecommerce"
)

const (
	UserPort  = "9001"
	EventPort = "9002"
)

var (
	Token = ""
)

func main() {
	testUserService()
	testEventService()
}

func testUserService() {
	// 拨号连接
	conn, err := grpc.Dial(":"+UserPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
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

// TokenAuth 通过实现 gRPC 中默认定义的 PerRPCCredentials，提供用于自定义认证的接口，它的作用是将所需的安全认证信息添加到每个 RPC 方法的上下文中。
type TokenAuth struct {
	token string
}

// GetRequestMetadata 获取当前请求认证所需的元数据
func (auth *TokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"token": auth.token}, nil
}

// RequireTransportSecurity 是否需要基于 TLS 认证进行安全传输
func (auth *TokenAuth) RequireTransportSecurity() bool { return false }

func testEventService() {
	// 拨号连接
	conn, err := grpc.Dial(":"+EventPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(&TokenAuth{
			token: "0",
		}),
	)
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

	// 访问该函数，获取 响应

	// 创建新事件
	resp1, err := client.CreateEvent(ctx, &pb.CreateEventRequest{
		Item: &ecommerce.Event{
			Name:      "test1",
			Content:   "test Event function",
			EndTimeAt: time.Now().Add(3 * time.Hour).Unix(),
			Type:      1,
			State:     1,
		},
	})
	if err != nil {
		log.Fatalf("client.CreateEvent err: %v", err)
	} else {
		log.Printf("Greeting: %s", resp1.String())
	}

	resp4, err := client.ShowEvent(ctx, &pb.ShowEventRequest{
		Id: "1",
	})
	if err != nil {
		log.Fatalf("client.ShowEvent err: %v", err)
	}
	log.Printf("Greeting: %s", resp4.String())
}

type Event struct {
	Name      string `json:"name"`
	Content   string `json:"content"`
	EndTimeAt int64  `json:"endTime_at"`
	State     int    `json:"state,omitempty"`
	Type      int    `json:"type,omitempty"`
}
