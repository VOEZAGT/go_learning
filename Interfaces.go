package main
import "fmt"

func main(){
	str := "Golang is good!"
	strPrt := &str
	strPrtPrt := &strPrt

	fmt.Printf("str type is %T,and values is %v\n", str, str)
	fmt.Printf("str type is %T,and values is %v\n", strPrt, strPrt)
	fmt.Printf("str type is %T,and values is %v\n",strPrtPrt,*strPrtPrt)
}