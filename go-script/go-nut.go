package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type UserCredential struct {
	Username string
	Password string
}

var ValidCredential = UserCredential{
	Username: "a",
	Password: "b",
}

var InvalidCredential = UserCredential{
	Username: "c",
	Password: "d",
}

const (
	DebugIsOn       = true
	TestingModeIsOn = true
)

var (
	proxyURLString      = "http://localhost:8080"
	TargetUrl           = "http://localhost:1111"
	FileURLs            = "urls.txt"
	FileURLsForTesting  = "test-urls.txt"
	FilePaths           = "paths.txt"
	FilePathsForTesting = "test-paths.txt"
)

func main() {

	proxyUrl, _ := url.Parse(proxyURLString)

	// creating reusable HTTP client
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}

	// urls.txt vs test-url.txt
	fileURLToOpen := FileURLs
	if TestingModeIsOn {
		fileURLToOpen = FileURLsForTesting
	}

	// reading file input line by line to invoke HTTP request
	file, _ := os.Open(fileURLToOpen)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		invokeRequest(client, url, &ValidCredential)
		invokeRequest(client, url, &InvalidCredential)
		invokeRequest(client, url, nil)
	}
}

func invokeRequest(client *http.Client, url string, userCredential *UserCredential) {
	req, err := http.NewRequest("GET", url, nil)
	if userCredential != nil {
		req.SetBasicAuth(userCredential.Username, userCredential.Password)
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("there was an error \n", err)
	}
	defer resp.Body.Close()

	if DebugIsOn {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
		fmt.Println("Done")
	}
}
