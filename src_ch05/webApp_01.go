package main
import ("fmt"
		"net/http")
// respose and request data type
func index_handler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf (w, "Oh, Go is neat!")
}

func main () {
	// Take the path and call func
	http.HandleFunc ("/", index_handler)
	// localhost:8000
	http.ListenAndServe (":8000", nil)
}