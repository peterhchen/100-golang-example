package main

import ("fmt")

// (input-type) return-type
func add3 (x, y float64) float64 {
	return x + y
}

func main () {
	num1, num2 := 1.0, 2.0
	fmt.Printf (" add3 (%.2f, %.2f) = %.2f\n", num1, num2, add3(num1, num2))
}