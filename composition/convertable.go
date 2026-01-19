package main

import "fmt"

type Convertable struct{
	Engine
	Transmission
	StreeringWheel
}

func (c Convertable) ConvertTop(){
	fmt.Println("converting top...")
}