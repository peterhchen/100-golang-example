package main
/* <stiemapindex>
<sitemap> or <url>
<loc>https://www.washingtonpost.com/business/color-of-money-live-december-20/2018/12/13/5efd82f2-fe36-11e8-a17e-162b712e8fc2_livediscussion.html</loc>
<news:news>
</news:news>
<news:title>
Security worries hobble ambitions of China tech giant Huawei
</news:title>
*/
import ("fmt"
"net/http"
"io/ioutil"
"encoding/xml")
type SitemapIndex struct {
	Locations []string `xml:"url>loc"`
}
type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}
type NewsMap struct {
	Keyword string
	Location string
}
func main () {
	var s SitemapIndex
	var n News
	news_map := make (map [string]NewsMap)
	resp, _ := http.Get ("https://www.washingtonpost.com/news-business-sitemap.xml")
	bytes, _ := ioutil.ReadAll (resp.Body)
	xml.Unmarshal (bytes, &s)
	fmt.Printf ("XML Parser:")
	i := 0
	// Range return two values: _ is index, Location is value
	for _, Location := range s.Locations {
		fmt.Printf ("\n%s", Location)
		resp, _ := http.Get (Location)
		bytes, _ := ioutil.ReadAll (resp.Body)
		xml.Unmarshal (bytes, &n)
		for idx, _ := range n.Titles {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
			fmt.Printf ("\n%s", n)
		}
		i++
		if i > 2 {
			break
		}
	}
	for idx, data := range news_map {
		fmt.Println ("\n\n\n", idx)
		fmt.Println ("\n", data.Keyword)
		fmt.Println ("\n", data.Location)
	}
}