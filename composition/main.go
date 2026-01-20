package main

type startable interface{
	Start()
}
func enginerStarter(vehicles... startable){
	for _, c := range vehicles{
		c.Start()
	}
}

func main() {
	miniCooper := Convertable{Engine{name: "Mini Copper"}, Transmission{}, StreeringWheel{}}
	cyberTruck := Truck{Engine{"Tesla"}, Transmission{}, StreeringWheel{}, }

	miniCooper.Start()
	// cyberTruck.Start()		// started
	cyberTruck.engine.Start()	// started

	enginerStarter(miniCooper, cyberTruck)
}