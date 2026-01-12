package main

import "fmt"


var ErrFirstError = CustomError{Message: "First error"}

type CustomError struct{
	Message string
}

func (c CustomError) Error()string{
	return c.Message
}

func someFunction()error{
	return ErrFirstError
}

func main() {

	err := someFunction()
	fmt.Println(err)		// Println function calls err.Error() dynamically and this happens only for error interface
	
}