package main

import (
	"fmt"
	"reflect"
)

func main(){

	myMap := make(map[string]string)
	myMap["a"] = "1"

	myArray := [1]int{1}

	fmt.Println(reflect.TypeOf(myMap))
	fmt.Println(reflect.ValueOf(myMap))

	fmt.Println(reflect.TypeOf(myArray))
	fmt.Println(reflect.ValueOf(myArray))

	s := myStruct{"lushijie"}
	k := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	for i := 0; i < k.NumField(); i++ {
		fmt.Println(k.Field(i).Name)
	}
	for i := 0; i < k.NumMethod(); i++ {
		fmt.Println(k.Method(i).Name)
	}
	name := v.MethodByName("GetName").Call([]reflect.Value{})
	fmt.Println(name)
}

type myStruct struct {
	Name string
}

func (m myStruct) GetName() string{
	return m.Name
}

func (m myStruct) SetName(name string) {
	m.Name = name
}
