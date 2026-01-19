package main

import "fmt"

type Engine struct{
	name string
}

func (e Engine) Start(){
	fmt.Printf("%v engine starting...\n", e.name)
}

func (e Engine) Stop(){
	fmt.Printf("%v engine stopping...\n", e.name)
}