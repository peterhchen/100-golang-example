package main
import ("fmt"
	"net/http"
	"io/ioutil")

func main () {
	// make a request and get a response and render. Error store in _
	//resp, _ := http.Get ("https://www.washingtonpost.com/news-sitemap-index.xml")
	resp, _ := http.Get ("https://www.washingtonpost.com/news-business-sitemap.xml")
	bytes, _ := ioutil.ReadAll (resp.Body)
	string_body := string (bytes)
	fmt.Println (string_body)
	resp.Body.Close()
}