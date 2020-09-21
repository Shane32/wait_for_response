package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"io/ioutil"
)

func main() {
	var url = flag.String("url", "http://localhost/", "URL to poll")
	var responseCode = flag.Int("code", 200, "Response code to wait for")
	var timeout = flag.Int("timeout", 2000, "Timeout before giving up in ms")
	var interval = flag.Int("interval", 200, "Interval between polling in ms")
	var localhost = flag.String("localhost", "", "Ip address to use for localhost")
	var sleep = flag.Int("sleep", 0, "Duration to sleep prior to request in ms")
	var lookfor = flag.String("lookfor", "", "String to look for")
	flag.Parse()

	fmt.Printf("Version 1.3.0\n")
	if *sleep > 0 {
		fmt.Printf("Sleeping for %d ms\n", *sleep)
		sleepDuration2 := time.Duration(*sleep) * time.Millisecond
		time.Sleep(sleepDuration2)
	}
	fmt.Printf("Polling URL `%s` for response code %d for up to %d ms at %d ms intervals\n", *url, *responseCode, *timeout, *interval)
	if *lookfor!="" {
		fmt.Printf("Looking for `%s`\n", *lookfor)
	}
	startTime := time.Now()
	timeoutDuration := time.Duration(*timeout) * time.Millisecond
	sleepDuration := time.Duration(*interval) * time.Millisecond

	if *localhost!="" && strings.Contains(*url, "localhost") {
		*url = strings.ReplaceAll(*url, "localhost", *localhost)
	}
	for {
		res, err := http.Get(*url)
		if err == nil && res.StatusCode == *responseCode {
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			bodyString := string(body)
			if *lookfor=="" || *lookfor==*bodyString {
				fmt.Printf("Response header: %v\n", res)
				fmt.Printf("Response body: %s\n", *bodyString)
				os.Exit(0)
			}
		}
		time.Sleep(sleepDuration)
		elapsed := time.Now().Sub(startTime)
		if elapsed > timeoutDuration {
			fmt.Printf("Timed out\n")
			os.Exit(1)
		}
	}
}
