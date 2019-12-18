package src

import (
	"fmt"
	"seckill/common/util"
	"time"
)

func TestChannel()  {
	ball := make(chan map[string]string, 10)
	passBall := func(player string) {
		for {
			preBall, err := <-ball
			if !err {
				panic("fail");
			}
			prePlayer := preBall["curr"]
			if ifShoot() {
				fmt.Println(player, "投篮", "\n")
				if ifHit() {
					score := getScore()
					fmt.Println( "-- 投篮命中,获得", score, "分\n")
				} else {
					fmt.Println("-- 未命中\n")
				}
				goto Next
			}
			if preBall["pre"] == "" {
				fmt.Print(prePlayer, "发球", "\n")
				goto Next
			} else {
				fmt.Println(prePlayer, " 把球传给了 ", player, "\n")
				goto Next
			}
			Next:
				time.Sleep(time.Second)
				currBall :=  map[string]string{"curr": player, "pre": prePlayer}
				ball <- currBall

		}
	}
	go passBall("1号")
	go passBall("2号")
	go passBall("3号")
	go passBall("4号")
	go passBall("5号")
	//开球
	cocah := map[string]string{"pre": "", "curr":"Coach"}
	ball <- cocah
	//preBall, err := <-ball
	//fmt.Println(preBall, err)
	var c chan bool // 一个零值nil通道
	<-c
}

func ifShoot() bool {
	return 1 == util.RandRange(0,3);
}

func ifHit() bool {
	return 1 == util.RandRange(0,2);
}

func getScore() int {
	return util.RandRange(2,4)
}

func TestChannel2()  {
	ch := make(chan int)
	genRandom := func() {
		for i := 0; i < 5; i++ {
			ch<- util.RandRange(0,100);
		}
		wg.Done()
	}

	getRandom := func() {
		for  i := 0; i < 5; i++ {
			random := <-ch
			fmt.Print("随机数：", random, "\n")
		}
		wg.Done()

	}
	wg.Add(2)
	go genRandom()
	go getRandom()
	wg.Wait()
}