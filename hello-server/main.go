package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	pb "xuexiangban_go/xxb-grpc-study/hello-server/proto"

	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
)

// 在这个结构体中重写sayhello方法
// 因为SayHello方法绑定在UnimplementedSayHelloServer结构体中，所以要先重写此结构体
type server struct {
	pb.UnimplementedSayHelloServer
}

// 这个自己重写的函数参数Request和Response都是指针，返回值是地址
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// 这个地方绑定了proto文件中的message结构体，返回的结构体，通过给req *pb.HelloRequest传参，来对此进行调用
	return &pb.HelloResponse{ResponseMsg: "hello" + " " + req.RequestName + " " + req.Age}, nil
}

func main() {
	// 开启端口，9090端口
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	//创建grpc服务，用grpc中的方法NewServer()，创建一个服务，并用一个变量去接收,返回值是一个地址
	grpcServer := grpc.NewServer()
	// 在grpc服务端中去注册我们自己编写的服务
	// 第一个参数是grpc的地址，即grpcServer
	// 第二个参数是传入被注册的结构体，一定要是引用注册，即将server结构体进行注册
	pb.RegisterSayHelloServer(grpcServer, &server{})

	// 启动服务
	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
