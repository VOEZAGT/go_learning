package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main(){
	/*
	golang中的下划线有3种用法：
	1.忽略返回值
	比如某个函数返回三个参数，但是只需要其中的俩个，另外一个参数可以忽略。
	ex: v1,v2,_ := function(.....)

	2.用在变量(特别是接口断言)
	如果定义了一个接口:
	type Foo interface{
		Say()
	}
	然后定义了一个结构体(struct)：
	type Dog struct {
	}
	然后希望在代码中判断Dog这个struct是否实现了Foo这个Interface
	var _Foo = Dog{}
	上面这个语句可以用来判断Dog是否实现了Foo，用作判断断言，如果Dog没有实现Foo，则会编译报错。

	3.用来import package
	假设在代码的import中这样引入package
	import _ "test/foo"
	这表示在执行本段代码之前会先调用test/foo中的初始化函数(init),这种使用方式仅让导入的包做
	初始化，而不使用包中的其他功能。

	例如我们定义了一个Foo Struct，然后对他进行初始化
	package foo
	import "fmt"
	type Foo struct{
		Id		int
		Name	string
	}
	func init() {
		f := &Foo{Id:123,Name: "abc"}
		fmt.Printf("init foo object:%v\n",f)
	}

	然后在main函数里面引入test/foo
	package main
	import (
		"fmt"
		_ "test/foo"
	)
	func main(){
			fmt.Printf("hello world\n")
	}
	运行结果如下：
	init foo object：&{123 abc}
	hello world
	可以看到：在main函数输出"hello world"之前就已经对foo对象进行初始化了。
	*/
	nums := []int{1,2,3,4,5}
	sum := 0
	for _, num := range nums {
	/*because we don't need the index,so we ignored it with the blank identifier.*/
		sum += num
	}
	fmt.Println("sum:",sum)

	for i, num := range nums{
		if num == 3{
			fmt.Println("index:", i)

		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana", "c": "chips"}

	for k,v := range kvs{
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:",k)
	}
	/*
	Golang遍历map时的key随机化问题以及解决方法
	Go的map在底层时用hashmap实现的；导致高效的hash函数肯定不是对
	key做顺序散列的，所以，在使用Go语言map过程中，key-value的插入顺序与
	遍历map时key的访问顺序时不相同的。

	1.通过range遍历map时，key的顺序被随机化
	在《Go maps in action》中，作者确认了该现象，并且是设计者故意为之。
	"使用范围循环在映射上迭代时，不会指定迭代顺序，也不能保证从一次迭代到下一次迭代的顺序相同。
	自Go 1以来，运行时随机化映射迭代顺序，因为程序员依赖于先前实现的稳定迭代顺序。"

	解决方法：然后解决办法方案就是把map里面的key保存到一个数组里面，然后使用排序功能进行排序，在对数组进行range，拿出key去取map里面的值
	*/

	/*integers := map[int]int{5:101, 10:102, 20:205, 30:502}
	keys := make([]int, 0, len(kvs))
	for _,values := range integers{
		id,_ := strconv.ParseInt(values["id"],0,64)
		keys = append(keys,int(id))
	}*/
	sorted_keys := make([]string,0)
	for k, _ := range kvs{
		sorted_keys = append(sorted_keys,k)
	}
	sort.Strings(sorted_keys)

	for _, k := range sorted_keys{
		fmt.Printf("k=%v, v=%v\n", k, kvs[k])

	}


	for i, c := range "go" {
		fmt.Println(i, c)
	}

	i, _ := strconv.ParseInt("23", 5, 64)
	println(i)
}