package main

import (
	"fmt"
	"reflect"
)

type qualification2 interface {
	TypeofGame(game string)
	DurationofCareer() int
}

type PLAYER2 struct {
	Name string
	Age int
	GAME string
	Position string
	level string
}

func (player *PLAYER2) TypeofGame(game string){
	fmt.Printf("Welcome to %v,%v\n",game,player.Name)
}

func (player *PLAYER2) DurationofCareer() int{
	fmt.Printf("You had played in valley for %v years",player.Age-13)
	return player.Age
}

func main(){
	var athletics qualification2 = &PLAYER2{
		Name:"Faker",
		Age:23,
		GAME:"League of Legends",
		Position:"MidLane",
		level:"Master",
	}
	var setNamestr string = "TypeofGame"
	var setNamestr2 string = "DurationofCareer"
	valueofathletics := reflect.ValueOf(athletics)
	//get the method named TypeofGame
	typeofgameMethod := valueofathletics.MethodByName( setNamestr )
	//call the method named TypeofGame
	args := []reflect.Value{reflect.ValueOf("League of Legends")}
	typeofgameMethod.Call(args)

	//get the method named DurationofCareer
	durationofcareerMethod := valueofathletics.MethodByName( setNamestr2)
	durationofcareerMethod.Call([]reflect.Value{})
}