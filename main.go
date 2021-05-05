package main

import (
	"fmt"
	"strings"

	"github.com/anaskhan96/soup"
)

const url = "https://xkcd.com/"

func main() {
	imageDownloader(url)

}

func imageDownloader(u string) {

	if !strings.HasSuffix(u, "#") {
		for {
			fmt.Printf("Downloading page from %v...\n", u)
			resp, err := soup.Get(u)

			if err != nil {
				panic(err)
			}
			html := soup.HTMLParse(resp)
			comicElem := html.Find("div", "id", "comic").FindAll("img")
			for _, img := range comicElem {
				comicURL := "http:" + img.Attrs()["src"]
				fmt.Printf("Downloading image from %v...\n", comicURL)
			}

			prevLink := html.Find("a", "rel", "prev")
			u = "http://xkcd.com" + prevLink.Attrs()["href"]
		}

	}

}
