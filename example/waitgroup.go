package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	// waitBySleep()
	// waitByChannel()
	waitByWG()
}

func waitBySleep(){
	for i := 0; i < 10; i++{
		go fmt.Printf("wait by sleep, i = %d \n", i)
	}
	time.Sleep(time.Second)
}

func waitByChannel(){
	c := make(chan bool, 10)
	for i := 0; i < 10; i++{
		go func(i int) {
			fmt.Printf("wait by channel, i = %d \n", i)
			c <- true
		}(i)
	}
	for i := 0; i < 10; i++{
		<- c
	}
}

func waitByWG(){
	w := sync.WaitGroup{}
	w.Add(10)
	for i := 0; i < 10; i++{
		go func(i int) {
			fmt.Printf("wait by wg, i = %d \n", i)
			w.Done()
		}(i)
	}
	w.Wait()

}