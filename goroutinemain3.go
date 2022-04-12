package main

import "fmt"

func main(){
	go func(name string){
		fmt.Println("Hello " + name)
	}("xuan")
	//主goroutine
}
