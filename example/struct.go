package main

import (
	"fmt"
	"reflect"
)

func main()  {
	// t := myType{Name: "lushijie"}
	// printMyType(&t)

	t := myTypeTag{Name: "test"}
	t1 := reflect.TypeOf(t)
	nickname := t1.Field(0).Tag.Get("json")
	fmt.Printf(nickname)
}


type myTypeTag struct {
	Name string `json:"nickname"`
}

type myType struct {
	Name string
}

func printMyType(t *myType)  {
	println(t.Name)
}