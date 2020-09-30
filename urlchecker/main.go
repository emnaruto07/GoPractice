package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	c := make(chan string)
	go CheckUrl(c)
	<-c
}

//CheckUrl to check url
func CheckUrl(c chan string) {
	var withHttp string

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		domain := strings.ToLower(sc.Text())
		//withHttps = "https://" + domain
		withHttp = "http://" + domain

		resp, err := http.Get(withHttp)
		if err != nil {
			continue
		}
		fmt.Println(withHttp, resp.StatusCode, http.StatusText(resp.StatusCode))

	}
	c <- withHttp
}
