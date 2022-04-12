package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int){
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i) //返回值为*int时，需要导入的是i的地址
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:",&i)
}