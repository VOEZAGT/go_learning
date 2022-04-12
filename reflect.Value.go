package main

import (
	"fmt"
	"reflect"
)

type job interface {
	Career(job string)
	DurationofCareer() int
}

type colleague struct {
	Name string
	Age int
	job string
}

func (person *colleague) Career(job string){
	fmt.Printf("Hello,your identification is %v,your job is %v\n", person.Name,person.job)
}

func (person *colleague) DurationofCareer() int {
	fmt.Printf("Your duration for the career for %v years", person.Age)
	return person.Age
}

func main(){
	person := &colleague{
		Name: "小包",
	}

	valueOfcolleague := reflect.ValueOf(person).Elem()
	valueofName := valueOfcolleague.FieldByName("Name")
	fmt.Printf("The value of valueofName is %v\n",valueofName)
	//判断字符是否可以设定变量值
	if valueofName.CanSet(){
		valueofName.Set(reflect.ValueOf("小张"))
	}

	fmt.Printf("person name is changed as %v\n",person.Name)

	var person2 job = &colleague{
		Name: "小姜",
		Age:23,
	}

	valueOfcolleague2 := reflect.TypeOf(person2)
	valueofName2,_ := valueOfcolleague2.MethodByName("DurationofCareer")
	//将person接收器放在参数的第一位
	args := []reflect.Value{reflect.ValueOf(person2)}
	result := valueofName2.Func.Call(args)
	fmt.Printf("%v\n",result)
	fmt.Printf("result of run method is %v\n",result[0])
}