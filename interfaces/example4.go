package main

import "fmt"

type StdLogger struct{}

type Logger interface{
	Log(string)
}

func (StdLogger)Log(word string){
	fmt.Println(word)
}

type Service struct{
	Logger
}

// Example 4️⃣ Interface with Struct Embedding
func example4() {
	s := Service{Logger: StdLogger{}}
	s.Log("Service started...")
	s.Logger.Log("Service Ended...")
}