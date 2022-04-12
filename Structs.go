package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) *person {

	p := person{name:name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob",20})
	fmt.Println(person{name:"Alice",age:30})
	fmt.Println(person{name:"Fred"})
	fmt.Println(&person{name:"Ann",age:40})
	fmt.Println(newPerson("Johnson"))

	s := person{name:"Sean",age:50}
	fmt.Println("person:",s.name)
	fmt.Println("age:",s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.name = "yangshuo"
	fmt.Println(sp.name)

}