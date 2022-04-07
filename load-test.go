package main

import (
	"fmt"
	"net/http"
)

func main() {
	go makeRequest("http://dev.olasunkanmi-olayinka.me/healthz")
	go makeRequest("http://dev.olasunkanmi-olayinka.me/healthz")
	makeRequest("http://dev.olasunkanmi-olayinka.me/healthz")
}

func makeRequest(url string) {
	for true {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(resp.Status)
		}
	}
}
