package main

import (
	"testing"
	"regexp"
	"fmt"
	"net/http/httptest"
	"net/http"
)

const (
	REGEX_OK = `^\[([\d,.]+s\])+([\w:\ ]+)(.*)(with\ )+(\[+\d+\])`
	REGEX_ERR = `^\[([\d,.]+s\])+([\w:\ ]+):(.*)$`
)

func TestRequest_MustReturn_StringWithElapsedTime(t *testing.T) {
	chn := make(chan string)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Test CRGO =)")
	}))

	defer ts.Close()
	go Request(ts.URL, chn)

	matched, err := regexp.MatchString(REGEX_OK, <-chn)
	if err != nil {
		t.Error(err)
	}

	if !matched {
		t.Fail()
	}
}

func TestRequest_MustReturn_StringWithError(t *testing.T) {
	chn := make(chan string)
	url := "localhost"

	go Request(url, chn)
	matched, err := regexp.MatchString(REGEX_ERR, <-chn)
	if err != nil {
		t.Error(err)
	}

	if !matched {
		t.Fail()
	}
}
