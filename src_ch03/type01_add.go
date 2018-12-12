package main

import ("fmt")

// (input-type) return-type
func add (x float64, y float64) float64 {
	return x + y
}

func main () {
	var num1 float64 = 1.0
	var num2 float64 = 2.0
	fmt.Printf (" add (%.2f, %.2f) = %.2f\n", num1, num2, add(num1, num2))
}