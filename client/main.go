package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	proto "seckill/proto"
)

func main()  {
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	// 创建新的客户端
	greeter := proto.NewProductService("go.micro.api.product", service.Client())

	// 调用greeter
	rsp, err := greeter.Detail(context.TODO(), &proto.DetailRequest{})
	if err != nil {
		fmt.Println(err)
	}

	// 打印响应请求
	fmt.Println(rsp.Total)
}
