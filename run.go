package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var v interface{}
	var n [10]int

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10;i++{
		temp := r.Intn(100)
		v = i

		if (temp % 2) == 0 {
			v="hello"
			n[i] = temp
		}

		if _,ok := v.(int);ok {
			fmt.Printf("the initial value is %d\n,the result is %v\n", temp, v)
		}else{
			fmt.Printf("the except value is %v\n", n)

		}
	}
}
