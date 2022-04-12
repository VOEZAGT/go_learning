package main

import "fmt"

func main(){
	//var a [3]int 				//定义三个整数的数组
	//var a [3]int = [3]int{1,3,5}//声明并且初始化固定长度的数组
	a := [...]int{2,5,3}

	fmt.Println(a[0])			//first
	fmt.Println(a[len(a)-1])    //last

	//打印索引和元素
	for i,v := range a{
		fmt.Printf("%d:%d\n",i,v)
	}

	//仅打印元素
	for _,v := range a{
		fmt.Printf("%v ",v)
	}
}
