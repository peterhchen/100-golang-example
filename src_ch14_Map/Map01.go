package main
import "fmt"
func main () {
	grades := make (map[string]float32)

	fmt.Println ("\nAssign map")
	grades["Timmy"] = 42
	grades["Jess"] = 92
	grades["Sam"] = 57
	fmt.Println (grades)

	fmt.Println ("\nIndex")
	TimsGrade := grades["Timmy"]
	fmt.Println(TimsGrade)

	fmt.Println ("\nDelete")
	delete(grades, "Timmy")
	fmt.Println (grades)

	fmt.Println ("\nKey - Value:")
	for k, v := range grades {
		fmt.Println (k , ": ", v)
	}
}