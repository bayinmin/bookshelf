package utils

import (
	"crypto/tls"
	"gonuts/config"
	"net/http"
	"net/url"
)

var (
	testingMode bool
	initialized bool
)

func SetTestingMode(isOn bool) {
	testingMode = isOn
	initialized = true
}

func requireTestingModeInit() {
	if !initialized {
		Error("HTTPHelper.go - testing mode is not set. Set testing mode through setTestingMode before using function in this file!")
	}
}

func IsTestingModeOn() bool {
	requireTestingModeInit()
	return testingMode
}

func GetUrlFileName() (name string) {
	// urls.txt vs test-urls.txt
	fileURLToOpen := config.FileURLs
	if IsTestingModeOn() {
		fileURLToOpen = config.FileURLsForTesting
	}
	return fileURLToOpen
}

func GetPathFileName() (name string) {
	// paths.txt vs test-paths.txt
	fileURLToOpen := config.FilePaths
	if IsTestingModeOn() {
		fileURLToOpen = config.FilePathsForTesting
	}
	return fileURLToOpen
}

func InitHTTPClient(proxyUrl *url.URL) *http.Client {
	// creating reusable HTTP client
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
