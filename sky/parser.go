package main

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"fmt"
	"strconv"
)

func main() {

	resp, err := http.Get("https://www.naver.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	// title 태그의 적힌 text를 출력
	title := doc.Find("title").Text()
	fmt.Println(title)

	// html tag의 id가 news_cast인 태그의 text를 출력
	news_cast := doc.Find("#news_cast").Text()
	fmt.Println(news_cast)

	// div중 3번째 있는 div가져와 text 출력
	div_three, _ := doc.Find("div").Eq(2).Html()
	fmt.Println(div_three)

	// 함수를 전달해 사용할 수도 있다. 모든 div의 번호와 내용을 출력
	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		fmt.Println(strconv.Itoa(i) + "번째 div" + s.Text())
	})

}
