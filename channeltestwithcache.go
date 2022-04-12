package main

import (
	"fmt"
	"time"
)

func consumer(ch chan int) {
	//线程休息10s之后再从channel中读取数据
	time.Sleep(time.Second * 10)
	<-ch
}
func main() {
	//创建一个长度为2的channel
	ch := make(chan int, 2)
	go consumer(ch)

	ch <- 0
	ch <- 1
	//发送数据不被堵塞
	fmt.Println("I am free")
	ch <- 2
	fmt.Println("I can not go there within 10s!")

	time.Sleep(time.Second)
}
