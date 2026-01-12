package main

import "fmt"

type Engine struct{
	HorsePower int
}

type Car struct{
	Model string
	Engine
}
func embedding() {

	myCar := Car{Model: "Tata", Engine: Engine{HorsePower:250 }}

	fmt.Println("Car model:", myCar.Model)
	fmt.Println("Car horse power:", myCar.Engine.HorsePower)
	
}