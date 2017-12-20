package main

// scape fortify taxonomy

import (
	"fmt"
	"log"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type webInfo struct {
	kindom string
	max    int
}

func main() {
	webInfos := []webInfo{
		{"Input+Validation+and+Representation", 25},
		{"API+Abuse", 13},
		{"Security+Features", 29},
		{"Time+and+State", 3},
		{"Errors", 2},
		{"Code+Quality", 13},
		{"Encapsulation", 13},
		{"Environment", 33},
	}
	urlTmpl := "https://vulncat.hpefod.com/en/weakness?kingdom=%s&po=%d"
Loop:
	for _, wi := range webInfos {
		kindom, err := url.QueryUnescape(wi.kindom)
		if err != nil {
			kindom = wi.kindom
		}
		fmt.Println(kindom)
		for i := 1; i <= wi.max; i++ {
			url := fmt.Sprintf(urlTmpl, wi.kindom, i)
			doc, err := goquery.NewDocument(url)
			if err != nil {
				log.Printf("ERROR: %v", err)
				break Loop
			}
			doc.Find("h1").Each(func(_ int, s *goquery.Selection) {
				s.Find("a").Each(func(_ int, s *goquery.Selection) {
					fmt.Println("\t" + s.Text())
				})
			})
		}
	}
}
