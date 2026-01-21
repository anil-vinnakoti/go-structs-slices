package main

func buffered() {
	bufferedChannel := make(chan int, 2)
	
	bufferedChannel <- 4		// does not block
	bufferedChannel <- 5		// does not block
	// bufferedChannel <- 6		// this causes dead lock as the capacity ofthe above buffered is 2
	
	anExample()
	println("main finished")
}

func anExample(){
	someChannel := make(chan int)
	go func(){
		someChannel <- 5
		someChannel <- 8
		close(someChannel)		// You must close a channel when iterating with range so the receiver knows when to stop.
	}()

	for value := range someChannel{
		println(value)
	}
}

// Buffered channel:
// - Send blocks only when buffer is full
// - Receive blocks only when buffer is empty