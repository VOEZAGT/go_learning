package main

import (
	"fmt"
	"reflect"
)

type qualification interface {
	TypeofGame(game string)
	DurationofCareer() int
}

type PLAYER struct {
	Name string
	Age int
	GAME string
	Position string
	level string
}

func (player *PLAYER) TypeofGame(game string){
	fmt.Printf("Welcome to %v,%v\n",game,player.Name)
}

func (player *PLAYER) DurationofCareer() int{
	fmt.Printf("You had played in valley for %v years",player.Age)
	return player.Age
}

func main(){
	TypeOfPLAYER := reflect.TypeOf(PLAYER{})

	//通过#NumField获取结构体字段的数量
	fmt.Println(TypeOfPLAYER.NumField())
	for i := 0;i < TypeOfPLAYER.NumField();i++{
		fmt.Printf("field name is %s,type is %s,kind is %s\n",
			TypeOfPLAYER.Field(i).Name,TypeOfPLAYER.Field(i).Type,TypeOfPLAYER.Field(i).Type.Kind())
	}
	fmt.Println("*****************************************")
	fmt.Println("")
	//根据字段名称获取结构体内的成员字段类型对象
	//类似于查询功能。
	nameField,_ := TypeOfPLAYER.FieldByName("Name")
	fmt.Printf("%v\n",nameField)
	fmt.Printf("field' name is %s,type is %s,kind is %s\n",nameField.Name,nameField.Type,nameField.Type.Kind())
	gameField,boolean := TypeOfPLAYER.FieldByName("GAME")
	fmt.Printf("%v\n",gameField)
	fmt.Printf("%v\n",boolean)
	fmt.Printf("field' name is %s,type is %s,kind is %s\n",gameField.Name,gameField.Type,gameField.Type.Kind())

	var QUA qualification  = &PLAYER{} //声明一个qualification的接口，并且用PLAYER进行接收。
	//获取接口qualification的类型对象
	typeofQUA := reflect.TypeOf(QUA)
	//打印qualification的方法类型和名称。
	for i :=0;i < typeofQUA.NumMethod();i++{
		fmt.Printf("method is %s,type is %s,kind is %s.\n",typeofQUA.Method(i).Name,typeofQUA.Method(i).Type,typeofQUA.Method(i).Type.Kind())
	}//打印接口方法内的类型
	fmt.Println("**************************************************************")
	method,_ :=typeofQUA.MethodByName("TypeofGame")
	fmt.Printf("method is %s,type is %s,kind is %s.\n",method.Name,method.Type,method.Type.Kind())

}
