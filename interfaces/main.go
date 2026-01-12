package main

import "fmt"

// example 1
type Speaker interface{
	Speak() string	// interface{} => any
}

type Person struct{}
func (p Person)Speak()string{
	return "person speaking..."
}

type Parrot struct{}
func (p Parrot)Speak()string{
	return "parrot speaking..."
}


func saySomething(s Speaker){
	fmt.Println(s.Speak())
}



func main() {
	// example 1
	saySomething(Person{})
	saySomething(Parrot{})

	// example 2
	example2()

	// example 3
	example3()

	// example 4
	example4()

	// example 5
	example5()
	   
}