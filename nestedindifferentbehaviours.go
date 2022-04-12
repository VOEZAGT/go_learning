package main

import "fmt"

//游泳特性
type swimming struct{
}

func (swim *swimming) swim(){
	fmt.Println("Swimming is my ability")
}
//飞行特性
type flying struct{
}

func (fly *flying) fly(){
	fmt.Println("flying is my ability")
}

//同时具备上述特性的话，需要内嵌
type WildDuck struct{
	swimming
	flying
}
type DomesticDuck struct{
	swimming
}

func main(){
	wild := WildDuck{}
	wild.fly()
	wild.swim()

	domestic := DomesticDuck{}
	domestic.swim()
}