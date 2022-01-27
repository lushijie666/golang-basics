package main

import "time"

func main() {
	for i :=0; i < 10; i++ {
		go println(i)
	}
	time.Sleep(time.Second)
}
