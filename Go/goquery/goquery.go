package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://qiita.com/TakanoriVega/items/6d7210147c289b45298a")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("span").Each(func(i int, s *goquery.Selection) {
		id, exists := s.Attr("id")
		if exists {
			fmt.Println(id)
		}
	})
}
