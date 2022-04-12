package main

import "fmt"

func main() {
	//1.切片的生成方法
	source := [...]int{1, 2, 3}
	sli := source[0:1]

	fmt.Printf("sli value is %v\n",sli)
	fmt.Printf("sli cap is %v\n",cap(sli))
	fmt.Printf("sli len is %v\n",len(sli))
	fmt.Printf("source value is %v\n",source)

	sli[0] = 4
	fmt.Printf("sli value is %v\n",sli)
	fmt.Printf("sli cap is %v\n",cap(sli))
	fmt.Printf("sli len is %v\n",len(sli))
	fmt.Printf("source value is %v\n",source)

	//2.动态创建切片
	//通过make函数动态创建切片
	//make([]T,size,cap)
	//T:成员类型，size：当前切片具备的长度；cap：切片容量
	sli = make([]int,2,4)
	fmt.Printf("sli value is %v\n",sli)
	fmt.Printf("sli len is %v\n",len(sli))
	fmt.Printf("sli cap is %v\n",cap(sli))

	//3.声明新的切片
	ex := []int{1,2,3}
	fmt.Printf("ex value is %v\n",ex)
	fmt.Printf("ex len is %v\n",len(ex))
	fmt.Printf("ex cap is %v\n",cap(ex))

	//切片的动态扩容
	arr1 := [...]int{1,2,3,4}
	arr2 := [...]int{1,2,3,4}

	sli1 := arr1[0:2]
	sli2 := arr2[2:4]
	//切片中的起始位影响切片的capacity，即容量。
	//举例子来说,sli1的size=2，cap=4；由于size小于cap；
	sli3 := arr2[3:4]
	fmt.Printf("the cap of sli3 is %v\n",cap(sli3))

	fmt.Printf("sli1 pointer is %p,len is %v,cap is %v,value is %v\n",&sli1,len(sli1),cap(sli1),sli1)
	fmt.Printf("sli1 pointer is %p,len is %v,cap is %v,value is %v\n",&sli2,len(sli2),cap(sli2),sli2)

	newSli1 := append(sli1,5)
	fmt.Printf("sli1 pointer is %p,len is %v,cap is %v,value is %v\n",&newSli1,len(newSli1),cap(newSli1),newSli1)
	fmt.Printf("source arr1 become %v\n",arr1)

	newSli2 := append(sli2,6)
	fmt.Printf("sli1 pointer is %p,len is %v,cap is %v,value is %v\n",&newSli2,len(newSli2),cap(newSli2),newSli2)
	fmt.Printf("source arr2 become %v\n",arr2)

	//容量足够的sli1直接将append添加的新元素覆盖到原有数组arr1上，而cap不够的话，需要做扩容操作
	//申请了新的底层数组，不会在数组的基础上进行操作。
	//如果原有数据可以添加新的元素，即切片只想的数组后面还有空间，但是切片自身的容量已经饱和，此时进行append操作，同样会进行扩容，申请新的内存空间。

	arr3 := [...]int{1,2,3,4}
	sli4 := arr3[0:2:2] // size:2,cap=2;切片中参数:begin：end：cap
	fmt.Printf("sli4 pointer is %p, len is %v, cap is %v,value is %v\n",&sli4,len(sli4),cap(sli4),sli4)

	newSli3 := append(sli4,5)
	fmt.Printf("sli4 pointer is %p, len is %v, cap is %v,value is %v\n",&newSli3,len(newSli3),cap(newSli3),newSli3)

	//为了方便切片的数据快速复制，使用copy
	//copy(destSli,scrSli,[]T)
}
