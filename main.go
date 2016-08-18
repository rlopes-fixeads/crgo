package main

import (
	"fmt"
	"time"
	"io/ioutil"
	"os"
	"net/http"
)

func Request(url string, chn chan <-string) {
	start := time.Now()
	resp, err := http.Get(url)
	secs := time.Since(start).Seconds()

	if err != nil {
		chn <- fmt.Sprintf("[%.2fs] Error: %s", secs, err.Error())
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	chn <- fmt.Sprintf("[%.2fs] to request with [%v] for [%v]", secs, len(body), url)
}

func main() {
	start := time.Now()
	chn := make(chan string)

	for _, url := range os.Args[1:] {
		go Request(url, chn)
	}

	for range os.Args[1:] {
		fmt.Println(<-chn)
	}

	fmt.Printf("[%.2fs] Request time.\n", time.Since(start).Seconds())
}