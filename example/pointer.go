package main

import (
	"fmt"
)

func main()  {
	str := "i am is lushijie"
	strpointer := &str
	strvalue := *&str

	fmt.Println(str)
	fmt.Println(strpointer)
	fmt.Println(strvalue)

	str = "i am is zhangsan"

	//strvalue = *&str
	fmt.Println(str)
	fmt.Println(strvalue)

}
