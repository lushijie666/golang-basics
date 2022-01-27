package main

import (
	"fmt"
	"math/rand"
)

func main(){

	c := make(chan int)
	go func() {
		println("hello from goroutine")
		c <- 1
	}()
	println(<- c)

	c1 := make(chan int, 10)
	go func() {
		for i :=0; i < 10; i++ {
			n := rand.Intn(10)
			fmt.Printf("c1 add value %d \n", n)
			c1 <- n
		}
		close(c1)
	}()
	for v := range c1 {
		fmt.Printf("c1 get value %d \n", v)
	}
}
