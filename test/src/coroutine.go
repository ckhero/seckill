package src

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"seckill/common/util"
	"sync"
	//_ "go.uber.org/automaxprocs"
)

var wg sync.WaitGroup

func loop()  {
	file, err := os.OpenFile("src.log", os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil{
		fmt.Print(err)
		return
	}
	defer file.Close()
	fileWiter := bufio.NewWriter(file)
	outputStr := "src\n"
	for i := 0; i < 100000; i++ {
		fileWiter.WriteString(outputStr)
	}
	wg.Done()
}

func TestCoroutine()  {
	start := util.UnixMilli()
	runtime.GOMAXPROCS(16)
	runtime.NumCPU()
	wg.Add(32)
	for i := 0; i < 32; i ++ {
		go loop()
	}
	wg.Wait()
	end := util.UnixMilli()
	fmt.Printf("cost time %v %v", end - start, runtime.NumCPU())
}