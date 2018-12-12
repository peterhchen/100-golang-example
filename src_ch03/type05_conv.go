package main

import ("fmt")

func main () {
	var a int = 1
	var b float64 = float64(a)	// Type convert from integer to float
	x := a 						// Type transfet from integer to x 
	fmt.Printf ("a = %d, b=%f, x = %d", a, b, x)
}