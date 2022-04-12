package main

import (
	"fmt"
)

type Person2 struct{
	Name string//姓名
	Birth string //生日
	id	int64
}

//指针类型，用于修改内部信息
func (person *Person2) changeName(name string){
	person.Name = name
}

//非指针类型,仅仅用于打印输出
func (person Person2) PrintInfo(){
	fmt.Printf("My name is %v,birth is %v,id is %v\n",person.Name,person.Birth,person.id)
	//Attempt to modify the id
	person.id = 3
}

func main() {
	//连续两个声明和申请
	p1:=new(Person2)
	p1.Name = "zhangsan"
	p1.Birth = "1992-09-12"
	p1.id = 1

	p2 := &Person2{
		"keller",
		"1999-09-22",
		2,
	}
	p1.PrintInfo()
	p2.PrintInfo()

	p1.changeName("aiqing")
	p1.PrintInfo()
}

