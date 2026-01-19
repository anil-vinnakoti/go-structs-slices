package main

import "fmt"

type Transmission struct{}

func (t Transmission) ShiftUp(){
	fmt.Println("shifting up...")
}

func (t Transmission) ShiftDown(){
	fmt.Println("shifting down...")
}