package utils

import (
	"crypto/tls"
	"gonuts/config"
	"io"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
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
				InsecureSkipVerify: true, // skip SSL verification for Burp Suite testing
			},
		},
	}
}

func ServeHTTPServer(port int) {
	Info("The server will run with TLS by default. Use https:// to access...")
	Info("Files are served from the 'serve' directory. For example, https://localhost:443/f/test.txt will serve the file 'serve/test.txt'")
	Info("File root URL is https://<host>:443/f/")
	Info("Please make sure that 443_server.crt and 443_server.key are in the config directory. Generate them using OpenSSL if you don't have them.")
	mux := http.NewServeMux()

	mux.HandleFunc("/f/", func(w http.ResponseWriter, r *http.Request) {
		PrintAccessLog(r)
		path := filepath.Join("serve", strings.TrimPrefix(r.URL.Path, "/f/"))
		http.ServeFile(w, r, path)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		output := PrintAccessLog(r)
		w.Write([]byte(output))
	})

	sport := strconv.Itoa(port)
	server := &http.Server{
		Addr:     ":" + sport,
		Handler:  mux,
		ErrorLog: log.New(io.Discard, "", 0), // disable default logging
	}

	Processing("The server is listening on port " + sport + "...\n")

	if err := server.ListenAndServeTLS("config/443_server.crt", "config/443_server.key"); err != nil {
		log.Fatalf("Failed to start HTTPS server: %v", err)
	}

}
