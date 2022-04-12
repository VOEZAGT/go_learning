package main

import (
	"fmt"
	"reflect"
)

//结构体的定义
/*
type structName struct{
	value1 valueType1
	value2 valueType2
}
*/

//test
type Person struct{
	Name string // 姓名
	Birth string // 生日
	ID int64 //身份证号
}

//结构体的实例化和初始化
//1.声明实例化
func main() {
	//声明方式1
	var p1 Person
	p1.Name = "王小二"
	p1.Birth = "1990-12-12"

	//声明方式2
	p2 := new(Person)
	p2.Name = "李老头"
	p2.Birth = "1993-12-13"

	//声明方式3
	p3 := &Person{}//视为一次new实例操作，同样将返回指针类型。
	p3.Name = "张老四"
	p3.Birth = "1995-09-30"

	//初始化
	//可以选择使用类似JSON的键值表示方式可以对结构体的字段进行填充。字段的初始化是可选的。依然使用Person
	p4 := Person{
		Name:"黄老板",
		Birth:"1990-02-22",
	}
	//当结构体内的所有字段都需要初始化，并且字段的填充顺序与他们在结构体内定义的顺序抑制时，可以将键值对中的键省略，代码如下所示
	p5 := Person{
		"爱情",
		"1990-00-00",
		5,
	}
	fmt.Printf("%v\n",reflect.TypeOf(p5))
	fmt.Printf("%v\n",reflect.TypeOf(p4))//返回main.Person，即Person结构体的指针类型，即*Person
	//p5的初始化并不是将键值一一对应填充给Person结构体，仅仅是按顺序填充给了Person结构体，而省略了其中的值。
	//初始化p5使用的是取值实例化
}