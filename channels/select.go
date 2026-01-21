package main

import (
	"fmt"
	"time"
)

func selectPractice() {
	channelOne := make(chan int)
	channelTwo := make(chan int)

	go func() {
		for i := range 3 {
			channelOne <- i
			time.Sleep(time.Second)
		}
		close(channelOne)
	}()

	go func() {
		for i := range 3 {
			channelTwo <- i
			time.Sleep(time.Second)
		}
		close(channelTwo)
	}()

	select {											// select lets a goroutine wait on multiple channel operations.
	case value := <-channelOne:
		fmt.Println("recieved from channelOne:", value)
	case value := <-channelTwo:
		fmt.Println("recieved from channelTwo:", value)
	}
	// Blocks until one case is ready
	// If multiple are ready → random choice
	// default makes it non-blocking

	// <------------------------------------------------------------------------------------------------->
	// we use for select to iterate multiple time to get all the available values from the multiple channels
	// its an infinte loop, causes deadlock when there's nothing to read from a channel
	for {		
		select {
		case value := <-channelOne:
			fmt.Println("recieved from channelOne:", value)
		case value := <-channelTwo:
			fmt.Println("recieved from channelTwo:", value)
		}
	}
}
