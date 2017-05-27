package crawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"os"
)

func Crawl() error {
	doc, err := goquery.NewDocument("http://www.ise.shibaura-it.ac.jp")
	if err != nil {
		fmt.Println("url scraping failed")
	}

	res, err := doc.Html()
	if err != nil {
		fmt.Println("dom get failed")
	}

	ioutil.WriteFile("./sample.html", []byte(res), os.ModePerm)
	return nil
}
