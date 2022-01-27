package main

import "fmt"

func main()  {

	// 变长参数
	total := sum(1 ,2, 3)
	fmt.Printf("sum is %d \n", total)

	// 多值返回
	result1, total1 := sum1(1 ,2)
	fmt.Printf("result is %s, sum is %d \n", result1, total1)

	// 命名返回值
	result2, total2 := sum2(2 ,3)
	fmt.Printf("result is %s, sum is %d \n", result2, total2)

	// 忽略返回值
	_, total3 := sum3(3 ,4)
	fmt.Printf("sum is %d \n", total3)

	// 回调函数
	operation(1, increase)
	operation(2, decrease)
}

func sum(elements ... int) (int) {
	var a = 0
	for _, v := range elements{
		a += v
	}
	return a
}

func sum1(a int, b int) (string, int){
	return "success", a + b
}

func sum2(a int, b int) (result string, total int) {
	result = "success"
	total = a + b
	return result, total
}

func sum3(a int, b int) (string, int) {
	return "success", a + b
}

func increase(a, b int) {
	 fmt.Printf("increase result is %d \n", a + b)
}

func decrease(a, b int) {
	fmt.Printf("decrease result is %d \n", a - b)
}

func operation(v int, f func(int, int)) {
	f(v, 1)
}
