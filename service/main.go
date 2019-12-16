package main

import (
	"fmt"
	"seckill/common/component"
	"seckill/common/constant"
	"seckill/common/util"
	"seckill/config"
)

func init()  {
	config.InitConfig(constant.GlobalService, util.GetArg(constant.ArgEnv, constant.GlobalDev))
	component.InitRedis()
}
func main() {
	fmt.Println(config.GetEnv())
	fmt.Println(config.AppConfig.Redis.Host)

}