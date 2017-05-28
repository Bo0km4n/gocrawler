package main

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"reflect"
)

func main() {
	doc, err := goquery.NewDocument("http://www.ise.shibaura-it.ac.jp")
	if err != nil {
		fmt.Println("url scraping failed")
	}

	fmt.Println(reflect.TypeOf(doc))
	getLink(doc)
}

func getLink(doc *goquery.Document) ([]string, error) {
	var url_list []string
	fmt.Println("test")
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		url_list = append(url_list, url)
	})
	if len(url_list) > 256 {
		return url_list[0:255], errors.New("url_list is too long")
	}
	fmt.Println(reflect.TypeOf(url_list))
	return url_list, nil
}
