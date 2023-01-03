package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "xuexiangban_go/xxb-grpc-study/hello-server/proto"
)

func main() {
	// client端链接server端，使用Dial函数
	// Dial函数第一个参数传入链接地址，第二个参数传入对象：与安全校验有关
	// WithTransportCredentials是采用不加密
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 所有的链接用完后一定要关闭
	defer conn.Close()

	// 建立链接
	client := pb.NewSayHelloClient(conn)

	// 执行rpc调用（这个方法在服务端来实现并返回结果）
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "abc", Age: "19"})

	fmt.Println(resp.GetResponseMsg())

}
