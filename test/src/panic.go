package src

import (
	"fmt"
	"os"
)

func TestPanic()  {
	defer func() {
		v := recover()
		if (v != nil) {
			fmt.Println(v)
		}
	}()

	var str = "world"
	// 这里，转换[]byte(str)将不需要一个深复制。
	for i, b := range str {
		fmt.Println(i, ":", b)
	}
	testtest(1,1)
	//aa := map[interface{}]interface{}{1:"a","a":"3", "adsfadaf": 1}
	//for key, val:= range aa {
	//	fmt.Println("key => val =", key, " + ", val)
	//	aa["adsfadaf"] = "test"
	//	fmt.Println(aa)
	//}
	//
	//arr := [4]int{1,2,3,4}
	//for key, val:= range &arr {
	//	fmt.Println("key => val =", key, " + ", val)
	//	arr[3] = 5
	//	fmt.Println(aa)
	//}
	//
	//arr2 := []int{}
	//for key, val:= range arr2[:] {
	//	fmt.Println("key => val =", key, " + ", val)
	//	arr2[3] = 5
	//	fmt.Println(aa)
	//}
}

func testtest(int, int)  {
	fmt.Println(os.Args)
}