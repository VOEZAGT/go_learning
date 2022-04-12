package main

import (
	"bufio"
	"fmt"
	"os"
)

func printInput(ch chan string) {
	//使用for循环从channel中读取数据
	for val := range ch {
		//如果出现EOF，结束
		if val == "EOF" {
			break
		}
		fmt.Printf("Input is %s\n", val)
	}
}
func main() {
	//创建一个无缓冲的channel
	ch := make(chan string)
	go printInput(ch)

	//从命令行中读取输入
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		val := scanner.Text()
		ch <- val
		if val == "EOF" {
			fmt.Println("Game is over!")
			break
		}
	}
	//程序的最后需要关闭ch
	defer close(ch)
}
