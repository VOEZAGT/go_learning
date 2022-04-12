package main

import (
	"container/list"
	"fmt"
)

func main(){
	tmpList := list.New()

	for i:=1;i <= 10;i++ {
		tmpList.PushBack(i)
		//*list:PushBack inserts a new element e with value v at the back of list l and returns e.
	}

	first := tmpList.PushFront(0)
	fmt.Printf("the value of the first is %v\n",*first)
	tmpList.Remove(first)

	for l:= tmpList.Front(); l != nil;l = l.Next(){
		fmt.Print(l.Value," ")
	}
}