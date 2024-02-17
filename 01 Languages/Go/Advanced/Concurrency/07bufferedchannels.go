package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 3)
	go sell(ch, &wg)
	wg.Wait()
}

func sell(ch chan int, s *sync.WaitGroup) {
	ch <- 10
	ch <- 20
	ch <- 30
	go buy(ch, s)
	fmt.Println("Sent data to the channel")
	s.Done()
}

func buy(ch chan int, s *sync.WaitGroup) {
	fmt.Println("Waiting for data")
	fmt.Println("Received data:",<-ch)
	s.Done()

}