package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	// unsafeWrite()
	// safeWrite()
	go lock()
	//go rlock()
	//go wlock()
	time.Sleep(time.Second)
}

func unsafeWrite() {
	conflictMap := make(map[int]int, 1)
	for i := 0; i < 10; i++ {
		go func() {
			conflictMap[0] = i
		}()
	}
}

func safeWrite() {
	s := SafeMap{
		safeMap: map[int]int{},
		Mutex: sync.Mutex{},
	}
	for i := 0; i < 10; i++ {
		go func() {
			s.Write(0, i)
		}()
	}
}

type SafeMap struct {
	safeMap map[int]int
	sync.Mutex
}

func (s *SafeMap)Read(k int) (int, bool){
	s.Lock()
	defer s.Unlock()
	result, ok := s.safeMap[k]
	return result, ok
}

func (s *SafeMap)Write(k,v int){
	s.Lock()
	defer s.Unlock()
	s.safeMap[k] = v
}

func lock(){
	s := sync.Mutex{}
	for i := 0; i < 10; i++ {
		s.Lock()
		defer s.Unlock()
		fmt.Printf("lock: %d \n", i)
	}
}

func rlock() {
	s := sync.RWMutex{}
	for i := 0; i < 10; i++ {
		s.RLock()
		defer s.RUnlock()
		fmt.Printf("rlock: %d \n", i)
	}
}

func wlock()  {
	s := sync.RWMutex{}
	for i := 0; i < 10; i++ {
		s.Lock()
		defer s.Unlock()
		fmt.Printf("wlock: %d \n", i)
	}
}

