package main

import (
	"fmt"
)

func test(){
	fmt.Println("I am work in a single goroutine")
}

func main(){
	go test()
	//主 goroutine 休眠 1 s
	//time.Sleep(time.Second)
}