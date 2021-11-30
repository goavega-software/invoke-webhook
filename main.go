package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := os.Getenv("URL")
	if url == "" {
		log.Fatal("url not defined")
	}
	method := os.Getenv("METHOD")
	if method == "" {
		method = "POST"
	}
	payload := os.Getenv("PAYLOAD")
	post(url, method, payload)
}

func post(url string, method string, payload string) {
	fmt.Println("{0} Performing Http {1}...", url, strings.ToUpper(method))
	var data []byte
	if (method == "post" || method == "put") && payload != "" {
		data = []byte(payload)
	}

	client := &http.Client{}
	var req *http.Request
	if data != nil {
		log.Print("With data")
		req, _ = http.NewRequest(method, url, bytes.NewBuffer(data))
	} else {
		req, _ = http.NewRequest(method, url, nil)
	}
	if os.Getenv("AUTH") != "" {
		req.Header.Set("Authorization", os.Getenv("AUTH"))
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	fmt.Println("status is", resp.StatusCode)
}
