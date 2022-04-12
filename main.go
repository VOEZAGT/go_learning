package main

import (
	"fmt"
	"reflect"
)


type Persons interface {
	hello(name string) string
	Run() string //return type is "string"
}

type Hero struct {
	Name string
	Age  int
	Speed int
}

func (hero *Hero) hello(name string) string{
	fmt.Println("Hello " + name + ",I am " + hero.Name)
	return hero.Name
}

func (hero *Hero) Run() string{ //notify the return type
	fmt.Printf("I'm running at %v km/h\n", hero.Speed)
	return "Running"
}

func main(){
	var people Hero //declare a struct named people
	people.Name = "Steven Liv"
	people.Age = 23
	people.Speed = 150
	fmt.Printf("Name is %s\n",people.Name)
	fmt.Printf("Age is %v\n",people.Age)
	fmt.Printf("Speed is %v\n",people.Speed)
	//获取实例的反射类型对象
	typeOfHero := reflect.TypeOf(Hero{})
	fmt.Printf("Hero's type is %s, kind is %s\n", typeOfHero, typeOfHero.Kind())
	fmt.Printf("*Hero's type is %s, kind is %s",reflect.TypeOf(&Hero{}),reflect.TypeOf(&Hero{}).Kind())
	/*
	第一个反射对象是Hero{}，指的是一个叫Hero的结构体；因此reflect.TypeOf(Hero{})返回的是变量所属的类型，即main.Hero;
	Kind指的是变量类型所归属的品种，在这里这的是struct。
	第二个反射对象是*Hero{},指向了类型为Hero的struct变量，即指针变量。因此说明&Hero{}的类型是*main.Hero,归属于种类ptr.
	对于指针类型的变量，可以使用Type.Elem获取到指针指向变量的真实类型对象。
	*/
	typeOfPtrHero := reflect.TypeOf(&Hero{})
	fmt.Printf("&Hero's type is %s, kind is %s\n",typeOfPtrHero,typeOfPtrHero.Kind())
	typeOfHero2 := typeOfPtrHero.Elem()
	fmt.Printf("typeOfPtrHero elem to typeOfHero, Hero's type is %s,kind is %s\n", typeOfHero2,typeOfHero2.Kind())
	/*通过typeOfPtrHero.Elem, 我们可以获取到*main.Hello指针指向变量的真实类型main.Hero的类型对象。*/

	//使用反射调用接口方法
	//声明Person接口，并且使用Hero结构体来实现。
	var person Persons = &Hero{
		Name:"AlinaSylas",
		Speed: 155,
	}
	var setNameStr string = "hello"
	valueOfPerson := reflect.ValueOf(person)
	fmt.Printf("%v\n",valueOfPerson)
	//return &{Alina Sylas 23 155}
	//get method named hello
	HelloMethod := valueOfPerson.MethodByName( setNameStr )
	//build the parameters and method-call via #CALL
	//HelloMethod.Call([]reflect.Value{reflect.ValueOf("校长")})//有点疑惑
	args := []reflect.Value{ reflect.ValueOf("Mike")  }
	fmt.Println(args)
	HelloMethod.Call(args)
	//get method named Run
	RunMethod := valueOfPerson.MethodByName("Run")
	result := RunMethod.Call([]reflect.Value{})
	fmt.Printf("the type of result is %v,result of run method is %s.\n", reflect.TypeOf(result).Kind(),result[0])

}

