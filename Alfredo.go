package main

import (
	"fmt"
	"math"
)

type pizza struct {
	numToppings float64
	diameter    float64
}

var bambini = pizza{6, 14}
var famiglia = pizza{2, 32}

const topCost = 0.002
const doughCost = 0.001
const margin = 1.6

func main() {
	bamPrice := alfredo(bambini)
	famPrice := alfredo(famiglia)

	//println(bambini) figure out to string
	fmt.Println(bamPrice)
	//println(famiglia)
	fmt.Println(famPrice) //If you don't use fmt it will show sci notation

	fmt.Println(pizzaCompare())
}

func pizzaCompare() bool {
	return alfredo(bambini) > alfredo(famiglia)
}

func alfredo(p pizza) float64 {
	area := math.Pi * math.Pow(p.diameter/2, 2)

	return (area*doughCost + area*p.numToppings*topCost) * margin
}
