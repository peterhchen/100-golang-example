package main

import ("fmt")

func main () {
	x := 15
	a := &x
	fmt.Printf ("a = %x\n", a)
	fmt.Printf ("*a = %d", *a)
}