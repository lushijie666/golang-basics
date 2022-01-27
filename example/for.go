package main

import "fmt"

func main()  {
	// 普通循环
	for i := 0; i < 10; i ++ {

	}

	// while循环等价
	a := 1
	for ; a < 10; {
		fmt.Println(a)
	}

	// 无限循环
	for {
		fmt.Println("hello world")
	}

	// 遍历字符串
	myStr := "hello world"
	for index, value := range myStr {
		fmt.Printf("index is %d, value is %d", index, value)
	}

	// 遍历数组
	myArray := []int{1, 2, 3}
	for index, value := range myArray {
		fmt.Printf("index is %d, value is %d", index, value)
	}

	// 遍历map
	myMap := make(map[int]int)
	for key, value := range myMap {
		fmt.Printf("key is %d, value is %d", key, value)
	}
}
