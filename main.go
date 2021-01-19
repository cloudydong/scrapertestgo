package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

var url = "https://finance.naver.com/item/coinfo.nhn?code=005930"

func main() {
	c := colly.NewCollector()
	c.OnHTML("body", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
		f, err := os.Create("corp.txt")
		errCheck(err)
		defer f.Close()
		_, fErr := f.WriteString(h.Text)
		errCheck(fErr)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.Visit(url)
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
