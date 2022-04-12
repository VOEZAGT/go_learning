package main

import "fmt"

func main(){
	classMates1 := make(map[int]string)

	//添加映射关系
	classMates1[0] = "xiaoming"
	classMates1[1] = "xiaohong"
	classMates1[2] = "xiaozhang"

	//根据key获取value
	fmt.Printf("id %v is %v\n",1,classMates1[1])

	//在声明时初始化数据
	classMates2 := map[int]string {
		0:"xiaofang",
		1:"xiaokai",
		2:"xiaozhang",
	}
	fmt.Printf("id %v is %v\n",3,classMates2[3])

}
