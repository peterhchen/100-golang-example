package main
import ("fmt"
	"net/http")

func index_handler (w http.ResponseWriter, r *http.Request) {	
	fmt.Fprintf (w, `<h1>Hey There</h1>
<h2>Go is fast!</h2>
<p>Go is simple!</p>
Hello World`)
}

func main () {
	http.HandleFunc ("/", index_handler)
	http.ListenAndServe (":8000", nil)
}