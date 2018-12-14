package main
/*
<stiemapindex>
<sitemap> or <url>
<loc>https://www.washingtonpost.com/business/color-of-money-live-december-20/2018/12/13/5efd82f2-fe36-11e8-a17e-162b712e8fc2_livediscussion.html</loc>
</sitemap> or </url>
<sitemap> or <url>
<loc>https://www.washingtonpost.com/business/boy-scouts-exploring-all-options-to-address-fiscal-woes/2018/12/13/541f7ca2-ff10-11e8-a17e-162b712e8fc2_story.html</loc>
</sitemap> or </url>
</sitemapindex>
*/
import ("fmt"
"net/http"
"io/ioutil"
"encoding/xml")
type SitemapIndex struct {
/*
[5] type == array
[5][5] int == array
[] type == slice
*/
	//Locations []Location `xml:"sitemap"`
	Locations []Location `xml:"url"`
}
type Location struct {
	Loc string `xml:"loc"`
}

func main () {
	// make a request and get a response and render. Error store in _
	//resp, _ := http.Get ("https://www.washingtonpost.com/news-sitemap-index.xml")
	resp, _ := http.Get ("https://www.washingtonpost.com/news-business-sitemap.xml")
	bytes, _ := ioutil.ReadAll (resp.Body)
	resp.Body.Close()
	var s SitemapIndex
	xml.Unmarshal (bytes, &s)
	fmt.Println (s.Locations)
}