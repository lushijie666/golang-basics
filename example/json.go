package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := Human{"lushijie"}
	j := marshal2JsonString(s)
	fmt.Println(j)
	a := unmarshal2struct(j)
	fmt.Println(a)
}

func unmarshal2struct(str string) Human {
	h := Human{}
	err := json.Unmarshal([]byte(str), &h)
	if err != nil {
		println(err)
	}
	return h
}

func marshal2JsonString(o Human) string {
	bytes, err := json.Marshal(&o)
	if err != nil {
		println(err)
	}
	return string(bytes)
}

type Human struct {
	Name string
}
