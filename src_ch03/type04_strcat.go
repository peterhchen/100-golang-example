package main

import ("fmt")

// (input-type) return-type
func multipleString (x, y string) (string, string) {
	return x, y
}

func main () {
	w1, w2 := "Hello", "There"
	fmt.Println (multipleString(w1, w2))
}