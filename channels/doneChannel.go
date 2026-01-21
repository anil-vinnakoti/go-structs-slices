package main

import (
	"fmt"
	"sync"
	"time"
)

func doneChannel() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	
	// dead lock example
	// doSomeWork(&wg)

	// done channel example
	doSomeWorkWithDoneChannel(done, &wg)
	time.Sleep(time.Second)
	close(done)

	wg.Wait()

}

func doSomeWork(wg *sync.WaitGroup) <-chan int {
	channel := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		channel <- 4									// causes dead lock as there is no reader for this channel
		fmt.Println("go routine finished")
	}()

	// fmt.Println("printing ch value:", <-channel)		// this make works fine as we are reading from the channel
	return channel
}

func doSomeWorkWithDoneChannel(done chan struct{}, wg *sync.WaitGroup) <-chan int {
	channel := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		select{
		case channel <- 5:
			fmt.Println("insertion of 5 done into channel ")
		case <-done:
			fmt.Println("closing channel")
		}
		
		fmt.Println("go routine finished")
	}()

	return channel
}
