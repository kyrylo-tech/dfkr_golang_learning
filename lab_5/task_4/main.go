package main

import (
	"task_4/market"
)

func main() {
	m := market.NewMarket()

	m.AddProduct("PC #1", 100, 20)
	m.AddProduct("PC #2", 150, 25)
	m.AddProduct("Microphone", 60, 6)
	m.AddProduct("Monitor", 80, 7)
	m.AddProduct("Keyboard", 100, 10)

	for { 
		m.GetProductsDialog()
	}
}