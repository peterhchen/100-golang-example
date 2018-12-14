package main

import ("fmt"
"encoding/xml")
var washPostXML = []byte (`
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:news="http://www.google.com/schemas/sitemap-news/0.9">
<url>
	<loc>
		https://www.washingtonpost.com/business/color-of-money-live-december-20/2018/12/13/5efd82f2-fe36-11e8-a17e-162b712e8fc2_livediscussion.html
	</loc>
</url>
<url>
	<loc>
		https://www.washingtonpost.com/business/boy-scouts-exploring-all-options-to-address-fiscal-woes/2018/12/13/541f7ca2-ff10-11e8-a17e-162b712e8fc2_story.html
	</loc>
</url>
</urlset>`)
type SitemapIndex struct {
	//Locations []Location `xml:"sitemap"`
	Locations []Location `xml:"url"`
}
type Location struct {
	Loc string `xml:"loc"`
}
func main () {

	bytes := washPostXML
	//fmt.Println (bytes)
	var s SitemapIndex
	xml.Unmarshal (bytes, &s)
	fmt.Println (s.Locations)
}