package main

import (
	"fmt"
	"math"
)

// Calculate area and circumference
func Calculate(d float64) (float64, float64) {
	// calculate area
	area := math.Pi * math.Pow(d/2, 2)
	// calculate circumference
	circumference := math.Pi * d

	return area, circumference
}

func main() {
	a, c := Calculate(14)
	fmt.Printf("the area is : %v", a)
	fmt.Println()
	fmt.Printf("the circumference is : %v", c)
}
