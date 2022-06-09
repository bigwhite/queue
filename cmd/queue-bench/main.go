package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/bigwhite/queue"
)

const (
	/*
		producerCount  = 100000
		enqueCountPerP = 100
		consumerCount  = 10
		dequeCountPerC = 100
	*/
	producerCount  = 100
	enqueCountPerP = 1000
	consumerCount  = 10
)

func main() {
	q := queue.NewChanQueue()
	var wgp sync.WaitGroup
	var wgc sync.WaitGroup
	var start = make(chan struct{})
	wgp.Add(producerCount)
	wgc.Add(10)

	fmt.Println("create Ps and Cs...")

	// p
	for i := 0; i < producerCount; i++ {
		go func() {
			defer wgp.Done()
			<-start
			for i := 0; i < enqueCountPerP; i++ {
				q.Enqueue(1)
			}
		}()
	}

	// c
	for i := 0; i < consumerCount; i++ {
		go func() {
			defer wgc.Done()
			for i := 0; i < producerCount*enqueCountPerP/consumerCount; i++ {
				_ = q.Dequeue()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	t1 := time.Now()
	fmt.Println("go!")
	close(start) //start
	go func() {
		wgp.Wait()
		fmt.Printf("producer use %d milliseconds\n", time.Since(t1).Milliseconds())
	}()
	wgc.Wait()
	fmt.Printf("consumer use %d milliseconds\n", time.Since(t1).Milliseconds())
}
