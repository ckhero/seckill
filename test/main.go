package main

import (
	"fmt"
	"seckill/common/component"
)

func main()  {
	//src.TestCoroutine()
	//src.TestDefer()
	//src.TestDefer2()
	var age component.Age = 1
	age.Increase()
	fmt.Print(age)
}

