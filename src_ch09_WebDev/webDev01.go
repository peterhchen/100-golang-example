package main
import ("fmt"
	"net/http")

func index_handler (w http.ResponseWriter, r *http.Request) {	
	fmt.Fprintf (w, "<h1>Hey There</h1>")
	fmt.Fprintf (w, "<h2>Go is fast!</h2>")
	fmt.Fprintf (w, "<p>Go is simple!</p>")
	fmt.Fprintf (w, "Hello World")
	fmt.Fprintf (w, "<p>You %s add %s</p>", "can", "<strong>variables</strong>")
}

func main () {
	http.HandleFunc ("/", index_handler)
	http.ListenAndServe (":8000", nil)
}