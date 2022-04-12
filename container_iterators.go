package main

import "fmt"

func main() {
	//数组的遍历
	nums := [...]int{1,2,3,4,5,6,7,8}
	for k,v := range nums{
		fmt.Println(k, v, " ")
	}
	fmt.Println()

	//切片的遍历
	slis := []int{1,2,3,4,5,6,7,8}
	for k,v := range slis{
		fmt.Println(k,v, " ")
	}

	//字段的遍历
	tmpMap := map[int]string{
		0 : "xiaoming",
		1 : "xiaohong",
		2 : "xiaozhang",
	}

	for k, v := range tmpMap{
		//k is key, v is value
		fmt.Println(k, v, " ")
	}

}
