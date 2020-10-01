package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	c := make(chan string)
	for i := 0; i < 100; i++ {
		go CheckUrl(c)
	}
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
		defer resp.Body.Close()
		Body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			continue
		}
		fmt.Println(withHttp, resp.StatusCode, http.StatusText(resp.StatusCode), string(Body))

	}
	c <- withHttp
}
