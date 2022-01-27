package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	q := Queue{
		queue: []string{},
		cond: sync.NewCond(&sync.Mutex{}),
	}
	go func() {
		for {
			q.Enqueue("aaa")
			time.Sleep(time.Second * 2)
		}

	}()
	for {
		q.Dequeue()
		time.Sleep(time.Second)
	}

}


type Queue struct {
	queue []string
	cond *sync.Cond
}

func (q *Queue) Enqueue(item string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.queue = append(q.queue, item)
	fmt.Printf("<----enqueue append item = %s \n", item)
	q.cond.Broadcast()
}

func (q *Queue) Dequeue() {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if len(q.queue) == 0 {
		fmt.Println("dequeue wait")
		q.cond.Wait()
	}
	item := q.queue[0]
	fmt.Printf("---->dequeue get item = %s \n", item)
	q.queue = q.queue[1:]
}



