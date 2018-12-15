package main
/* <stiemapindex>
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
/* [5] type == array
[5][5] int == array
[] type == slice */
	Locations []Location `xml:"url"`
}
type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String () string {
	return fmt.Sprintf (l.Loc)
}
func main () {
	resp, _ := http.Get ("https://www.washingtonpost.com/news-business-sitemap.xml")
	bytes, _ := ioutil.ReadAll (resp.Body)
	resp.Body.Close()
	var s SitemapIndex
	xml.Unmarshal (bytes, &s)
	fmt.Printf ("XML Parser:")
	i := 0
	// Range return two values: _ is index, Location is value
	for _, Location := range s.Locations {
		fmt.Printf ("\n%s", Location)
		i++
		if i > 5 {
			break
		}
	}
}