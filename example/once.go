package main

import (
	"fmt"
	"sync"
)

func main(){
	o := sync.Once{}
	o.Do(func() {
		fmt.Println("first onece")
	})
	o.Do(func() {
		fmt.Println("second onece")
	})
	o.Do(func() {
		fmt.Println("third onece")
	})
}
