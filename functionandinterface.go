package main

import "fmt"

//结构提实现接口
//1.结构体实现接口
//结构体实现Invoker接口的代码如下：

//调用器接口
type Invoker interface{
	// 实现一个Call
	Call(interface{})
}

//结构体类型
type Struct struct{
}

//实现Invoker的Call
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

//函数体实现接口
type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}){
	f(p)
}


func main() {
	//声明接口变量
	var invoker Invoker

	//实例化结构体
	s := new(Struct)

	//将实例化的结构体赋值到接口
	invoker = s
	//使用接口调用实例化结构体的方法Struct.Call
	invoker.Call("hello")

	//将匿名函数函数转化为FuncCaller类型，再赋值给接口

	invoker = FuncCaller(func(v interface{}){
		fmt.Println("from function",v)
	})

	//使用接口调用FuncCaller.Call,内部会调用函数本体
	invoker.Call("hello")
}
