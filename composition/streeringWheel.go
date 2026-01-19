package main

import "fmt"

type StreeringWheel struct{}

func (sw StreeringWheel)TurningLeft(){
	fmt.Println("turning left...")
}

func (sw StreeringWheel)TurningRight(){
	fmt.Println("turning right...")
}