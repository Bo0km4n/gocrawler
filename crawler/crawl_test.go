package crawler

import (
	"fmt"
	"testing"
)

func TestCrawl(t *testing.T) {
	t_flag := Crawl()
	fmt.Println("crawl end")

	if t_flag != nil {
		t.Error("スクレイピングに失敗しています")
	}

}
