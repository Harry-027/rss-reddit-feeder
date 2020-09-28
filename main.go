package main

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/Harry-027/rss-reddit-feeder/models"
	"github.com/robfig/cron/v3"
)

func Crawler(address string) {
	resp, err := http.Get(address)
	if err != nil {
		fmt.Printf("Error GET: %v\n", err)
		return
	}
	defer resp.Body.Close()
	rss := models.Rss{}
	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		fmt.Printf("Error Decode: %v\n", err)
		return
	}
	fmt.Println("Category ::", rss.Category)
	fmt.Println("Updated ::", rss.Updated)
	fmt.Println("Title ::", rss.Title)
	fmt.Println("SubTitle ::", rss.Subtitle)
	fmt.Println("Entry ::", rss.Entry)
	for _, item := range rss.Entry {
		fmt.Println(item.Author)
		fmt.Println(item.Content)
	}
}

func WebCrawler() {
	for _, value := range models.Address {
		Crawler(value)
	}
}

func main() {
	WebCrawler()
	cr := cron.New()
	cr.AddFunc("1 * * * *", WebCrawler)
	cr.Start()
}
