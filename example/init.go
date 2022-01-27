package main

import "fmt"

var name = 0
func main()  {
	fmt.Printf("hello world, name is %d", name)
}

func init() {
	name = 1
}