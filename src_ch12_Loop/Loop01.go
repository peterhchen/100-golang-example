package main
import "fmt"

func main () {
	fmt.Println ("loop 1")
	for i := 0; i < 5; i++ {
		fmt.Println (i)
	}
	//while loop
	fmt.Println ("loop 2")
	fmt.Println()
	j := 0
	for j < 10 {
		fmt.Println (j)
		j+=5
	}
	fmt.Println ("loop 3")
	x := 5
	for {
		fmt.Println ("Do stuff", x)
		x += 3
		if x > 10 {
			break
		}
	}
	fmt.Println ("loop 4")
	for x := 5; x < 15; x += 3 {
		fmt.Println ("Do Stuff", x)
	}
}