package main



func main() {
	miniCooper := Convertable{Engine{name: "Mini Copper"}, Transmission{}, StreeringWheel{}}
	cyberTruck := Truck{Engine{"Tesla"}, Transmission{}, StreeringWheel{}, }

	miniCooper.Start()
	// cyberTruck.Start()		// started
	cyberTruck.engine.Start()	// started
}