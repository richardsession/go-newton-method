package main

import "fmt"

func Sqrt(x float64) float64 {
    var sqrt float64
    var temp float64 = x / 2   // Initial Seed
    var threshold float64 = 0.0000001
	
	for i := 0; i < 10; i++ {
		sqrt = temp - ((temp * temp - x)/(2 * temp))
		
		if (temp - sqrt) < threshold {
			return sqrt
		}
		
		temp = sqrt
	}
	
	return sqrt
}

func main() {
	fmt.Println(Sqrt(10))
}
