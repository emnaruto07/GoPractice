package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/anaskhan96/soup"
)

const url = "https://xkcd.com/2199"

var comicURL string

func main() {
	imageDownloader(url)

}

func imageDownloader(u string) {
	//Creating xkcd folder
	_, err := os.Stat("xkcd")
	if os.IsNotExist(err) {
		fmt.Println("[+]Created xkcd folder!")
		os.Mkdir("xkcd", 0755)
	} else {
		fmt.Println("[-]Removing old xkcd folder")
		os.RemoveAll(`xkcd`)
		fmt.Println("[+]Now Creating new xkcd Folder!!")
		os.Mkdir("xkcd", 0755)
	}
	//Checking if there is # in the end of the URL
	if !strings.HasSuffix(u, "#") {
		for {
			fmt.Printf("Downloading page from %v...\n", u)
			resp, err := soup.Get(u)
			if err != nil {
				panic(err)
			}
			html := soup.HTMLParse(resp)
			comicElem := html.Find("div", "id", "comic").FindAll("img")
			if len(comicElem) == 0 {
				fmt.Println("[!!]No image found!!")
				prevLink := html.Find("a", "rel", "prev")
				u = "http://xkcd.com" + prevLink.Attrs()["href"]
				continue
			}
			for _, img := range comicElem {
				comicURL = "http:" + img.Attrs()["src"]
				fmt.Printf("Downloading image from %v...\n", comicURL)
				resp, err := http.Get(comicURL)
				if err != nil {
					fmt.Println("[!!]Invalid URL")
				}
				//Save the image to ./xkcd
				os.Chdir("./xkcd")
				file, err := os.Create(comicURL[28:])
				if err != nil {
					log.Fatalln("[-]Check file names")
				}

				_, err = io.Copy(file, resp.Body)
				if err != nil {
					log.Fatal(err)
				}
				prevLink := html.Find("a", "rel", "prev")
				u = "http://xkcd.com" + prevLink.Attrs()["href"]
			}

		}

	}

}
