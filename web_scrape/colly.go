package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
	"github.com/gocolly/colly"
)
type article struct{
	Title string "json:'title'"
	Author string "json:'article'"
	Date string "json:'date_time'"
	Link string "json:'link'"
}
func extract_data(start,end int){
	result := make([]article,0)
	for i:=start;i<=end;i++{
		url := "https://www.ptt.cc/bbs/Lifeismoney/index"+strconv.Itoa(i)+".html"
		fmt.Println(url)
		c := colly.NewCollector()
		c.OnHTML(".r-ent", func(e *colly.HTMLElement) {
			title := e.ChildText(".title")
			author := e.ChildText(".author")
			datetime := e.ChildText(".date")
			link := e.ChildAttr("a","href")
			strings.Replace(title,"\n","",-1)		
			post := article{
				Title:title,
				Author:author,
				Date: datetime,
				Link: link,
			}
			fmt.Println(post)
			result=append(result, post)
		})
		c.OnRequest(func(r *colly.Request) {
			r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
		})
		c.Visit(url)
	}
	WriteJson(result)
}

func main() {
	var start int = 3460
	var end int =  3468
	extract_data(start,end)
}

func WriteJson(data []article){
	file,err := json.MarshalIndent(data,""," ")
	if err != nil{
		log.Println("Could not write data to json")
		return
	}
	_ = ioutil.WriteFile("test.json",file,0644)
}