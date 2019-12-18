package src

import "fmt"

func TestDefer()  {
	defer fmt.Println("9")
	fmt.Println("0")
	defer fmt.Println("8")
	fmt.Println("1")
	if true {
		defer fmt.Println("not reachable")
	}
	func() {
		defer fmt.Println("7")
		fmt.Println("3")
		defer func() {
			fmt.Println("5")
			fmt.Println("6")
		}()
		fmt.Println("4")
	}()
	fmt.Println("2")
	return
	defer fmt.Println("not reachable")
}

func TestDefer2()  {
	for i := 0; i < 3; i++ {
		defer func(i int) {
			// 此i为形参i，非实参循环变量i。
			fmt.Println("b:", i)
		}(i)
	}
	//fmt.Println(22)
	//func() {
	//	for i := 0; i < 3; i++ {
	//		defer func() {
	//			fmt.Println("b:", i)
	//		}()
	//	}
	//}()
}