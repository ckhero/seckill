package main

import (
	"fmt"
	micro "github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
	"seckill/common/component"
	"seckill/common/constant"
	"seckill/common/model"
	"seckill/common/util"
	"seckill/config"
	"seckill/service/handle"
	"sync"
	proto "seckill/proto"

)

var wg sync.WaitGroup
type ProductProto struct {}
func init()  {
	config.InitConfig(constant.GlobalService, util.GetArg(constant.ArgEnv, constant.GlobalDev))
	component.InitRedis()
	model.InitDB()
}
func main() {
	srv := micro.NewService(micro.Name("go.micro.api.product"))
	srv.Init()

	proto.RegisterProductHandler(srv.Server(), new(handle.Product))
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
	defer func() {

		if v := recover(); v != nil {
			logrus.Error(v)
		}

		component.Redis.Conn.Close()
		model.CloseDB()
	}()
}

func test() {
	//fmt.Println("aaa")
	//fmt.Println(time.Now().Unix())
	//wg.Done()
	component.Redis.SetStr("a2222", "a")
	component.Redis.SetStr("a22222222", "22a")
}