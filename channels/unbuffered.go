package main

func unbuffered() {
	var channel chan int
	// println( channel)		// gives the address of nil(0x0)
	// <-channel				// deadlock as it blocks the program
	// close(channel)			// closing nil channel causes panic


	channel = make(chan int)	// unbuffered channel
	go func(){
		channel <- 8			// this runs in a different env from main func(executed concurrently)
		channel <- 9			// this runs in a different env from main func(executed concurrently)
	}()

	// println(<- channel)			// this lines blocks the main func from exiting as the channel is expecting a value
	// println(<- channel)		// causes deadlock as we are only inserting only one value to the channel the in above goroutine 
}

//Unbuffered channel:
// - Send blocks until receive
// - Receive blocks until send