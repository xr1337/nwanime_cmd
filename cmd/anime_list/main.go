package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	result, err := GetWebBlock()
	if err != nil {
		panic(err)
	}
	fmt.Print(result)
}

// GetWebBlock - downloads the data from nwanime
func GetWebBlock() (text string, err error) {
	url := "https://www.nwanime.tv/home/"
	resp, err := http.Get(url)
	if err != nil {
		return text, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	data := Extract(resp.Body)
	return data, nil
}

// Extract the html contents for title-days
func Extract(r io.Reader) string {
	var str strings.Builder
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".video-item").Each(func(i int, s *goquery.Selection) {
		name := s.Find(".video-title span a")
		time := s.Find(".time")
		if len(name.Text()) > 0 {
			str.WriteString(fmt.Sprintf("%s - %s\n", name.Text(), time.Text()))
		}
	})
	return str.String()
}
