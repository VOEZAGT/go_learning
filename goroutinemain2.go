package main

import "fmt"

func setVTo1(v *int){
	*v = 1
}
func setVTo2(v *int){
	*v = 2
}

func main(){
	v := new(int)
	go setVTo1(v)
	go setVTo2(v)
	fmt.Println(*v)
}
