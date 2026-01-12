package main

import "fmt"

type Printer interface{
	Print(string)
}

type ConsolePrinter struct{}
func (ConsolePrinter)Print(msg string){
	fmt.Println(msg)
}

func log(p Printer){
	p.Print("Hello interface...")
}

// Example 2️⃣ Why Interfaces Are Useful (decoupling)
func example2() {
	log(ConsolePrinter{})
}
