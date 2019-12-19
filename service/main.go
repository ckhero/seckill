package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"seckill/common/component"
	"seckill/common/constant"
	"seckill/common/model"
	"seckill/common/util"
	"seckill/config"
	"sync"
)

var wg sync.WaitGroup

func init()  {
	config.InitConfig(constant.GlobalService, util.GetArg(constant.ArgEnv, constant.GlobalDev))
	component.InitRedis()
	model.InitDB()
}
func main() {
	//productExtend, err := new(model.ProductExtend).FindOne(2)
	//if productExtend != nil {
	//	fmt.Println("productExtend:", util.JsonEncode(productExtend))
	//}

	product, err := new(model.Product).FindOne(2)
	fmt.Println("product:", util.JsonEncode(product))
	productExtends, _ := product.GetProductExtends()
	fmt.Println("product:", util.JsonEncode(productExtends))
	if err != nil {
		fmt.Println(err)
	}
	logrus.Infof("done")
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