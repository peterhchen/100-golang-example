package main

import ("fmt")

func main () {
	x := 15
	a := &x
	fmt.Printf ("a = %x\n", a)
	fmt.Printf ("*a = %d\n", *a)
	
	*a = 5;
	fmt.Printf ("x = %x\n", x)
	*a = *a**a;
	fmt.Printf ("x = %d\n", x)
}