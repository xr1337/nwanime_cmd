package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const limit int = 14

type Anime struct {
	time  string
	title string
	url   string
}

func (a *Anime) String() string {
	return fmt.Sprintf("[%s]%s\n%s\n\n", a.time, a.title, a.url)
}

func main() {
	result, err := GetWebBlock()
	if err != nil {
		panic(err)
	}
	result = result[:limit]
	fmt.Print(AnimeFormat(result))
}

func AnimeFormat(animes []Anime) string {
	var str strings.Builder

	for _, anime := range animes {
		str.WriteString(anime.String())
	}
	return str.String()
}

// GetWebBlock - downloads the data from nwanime
func GetWebBlock() ([]Anime, error) {
	url := "https://www.nwanime.tv/home/"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	return Extract(resp.Body), nil
}

// Extract the html contents for title-days
func Extract(r io.Reader) []Anime {
	var items []Anime
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".video-item").Each(func(i int, s *goquery.Selection) {
		name := s.Find(".video-title span a")
		link := name.AttrOr("href", "")
		time := s.Find(".time")
		if len(name.Text()) > 0 {
			items = append(items, Anime{time.Text(), name.Text(), link})
		}
	})
	return items
}
