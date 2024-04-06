package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://youtube.com",
		"https://netflix.com",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkUrl(url, c)
	}

	for url := range c {
		go func(url string) {
			time.Sleep(time.Second)
			checkUrl(url, c)
		}(url)
	}

}

func checkUrl(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down!")
		c <- url
		return
	}
	fmt.Println(url, "is up!")
	c <- url
}
