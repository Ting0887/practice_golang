package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func HttpGet(url string)(result string,err error){
	 resp,err1 := http.Get(url)
	 if err1!=nil{
		err = err1
		return 
	 }
	 defer resp.Body.Close()

	 buf := make([]byte,4096)
	 for{
		 n, err2 := resp.Body.Read(buf)
		 if n == 0{
			 fmt.Println("success")
			 break
		 }
		 if err2 != nil && err2 != io.EOF{
			err = err2
			return
		 }
		 result += string(buf[:n])
	 }
	 return
}

func scrape(start,end int){
	fmt.Printf("爬取第%d頁到第%d頁", start,end)
	for i :=start;i<=end;i++{
		url := "https://www.ptt.cc/bbs/Baseball/index"+strconv.Itoa(i)+".html"
		fmt.Println(url)
		result,err := HttpGet(url)
		if err != nil{
			fmt.Println("Http Error:",err)
			continue
		}
		//fmt.Println("result:",result)
		f,err := os.Create("第"+strconv.Itoa(i)+"頁.html")
		f.Write([]byte(result))
		defer f.Close()
	}
}

func main(){
	var start,end int
	fmt.Println("請輸入啟示頁面")
	fmt.Scan(&start)
	fmt.Println("請輸入啟示頁面")
	fmt.Scan(&end)
	scrape(start,end)
}