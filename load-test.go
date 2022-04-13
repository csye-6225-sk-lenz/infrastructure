package main

import (
	"fmt"
	"net/http"
)

func main() {
	go makeRequest("http://prod.olasunkanmi-olayinka.me/healthtest3")
	go makeRequest("http://first@xampe.com:12345@prod.olasunkanmi-olayinka.me/v1/user/self")
	go makeRequest("http://prod.olasunkanmi-olayinka.me/healthz")
	go makeRequest("http://prod.olasunkanmi-olayinka.me/healthtest3")
	go makeRequest("http://first@xampe.com:12345@prod.olasunkanmi-olayinka.me/v1/user/self")
	makeRequest("http://prod.olasunkanmi-olayinka.me/healthz")
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
