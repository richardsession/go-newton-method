package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// Starting point
	var temp float64 = x / 2
	var z float64
	
	for i := 0; i < 10; i++ {
		z = temp - ((temp * temp - x)/(2 * temp))
		
		if (temp - z) < 0.0000001 {
			return z
		}
		
		temp = z
	}
	
	return z
}

func main() {
	fmt.Println(Sqrt(10))
}
