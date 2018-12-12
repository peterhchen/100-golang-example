package main

import ("fmt")

// (input-type) return-type
func add2 (x, y float64) float64 {
	return x + y
}

func main () {
	var num1, num2 float64 = 1.0, 2.0
	fmt.Printf (" add2 (%.2f, %.2f) = %.2f\n", num1, num2, add2(num1, num2))
}