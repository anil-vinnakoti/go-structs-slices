package main

import "fmt"

type engine interface{
	Start()
	Stop()
}

type Truck struct{
	engine				// with interface
	Transmission		// with struct
	StreeringWheel
}

func (t Truck)SwitchTo4WD(){
	fmt.Println("switching to 4 wheel drive...")
}