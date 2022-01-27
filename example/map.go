package main

import "fmt"

func main()  {
	myMap := make(map[string]string)
	myMap["a"] = "1"

	for key, value := range myMap {
		fmt.Printf("key is %s, value is %s \n", key, value)
	}

	value, exists := myMap["b"]
	if exists {
		fmt.Printf("b is exists, value is %s \n", value)
	} else {
		fmt.Printf("b is not exists, value is null \n")
	}

	myFuncMap := map[string]func() int{
		"funcA": func() int {
			return 1
		},
		"funcB" : func() int {
			return 2
		},
	}

	for key, value := range myFuncMap {
		fmt.Printf("key is %s, value is %s \n", key, value)
		f := myFuncMap[key]
		fmt.Printf("func result is %d \n", f())
	}
}
