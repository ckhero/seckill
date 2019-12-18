package main

import (
	"seckill/common/component"
	"seckill/common/constant"
	"seckill/common/util"
	"seckill/config"
	"sync"
)

var wg sync.WaitGroup

func init()  {
	config.InitConfig(constant.GlobalService, util.GetArg(constant.ArgEnv, constant.GlobalDev))
	component.InitRedis()
}
func main() {
	test()

	defer func() {
		component.Redis.Conn.Close()
	}()
}

func test() {
	//fmt.Println("aaa")
	//fmt.Println(time.Now().Unix())
	//wg.Done()
	component.Redis.SetStr("a2222", "a")
	component.Redis.SetStr("a22222222", "22a")
}