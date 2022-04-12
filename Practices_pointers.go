package main

import "fmt"

func main() {
	str := "Golang is Good!"
	strPrt := &str

	fmt.Printf("str type is %T,value is %v,address is %p\n\n", str, str, &str)
	fmt.Printf("str type is %T,value is %v,address is %p\n\n", strPrt, strPrt, &strPrt)

	newStr := &strPrt
	fmt.Printf("newStr type is %T,value is %v,address is %p\n\n", newStr, newStr, &newStr)

	*strPrt = "Java is Good too!"
	fmt.Printf("newStr type is %T,value is %v,address is %p\n\n",newStr,newStr,&newStr)
	fmt.Printf("str")

}