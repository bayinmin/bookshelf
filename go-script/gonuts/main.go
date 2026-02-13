package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"gonuts/attack"
	"gonuts/config"
	"gonuts/utils"

	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

var title = `
=========================================
== GO NUTs - The Go-based HTTP Fuzzer ===
=========================================
`

// bootstrap value
var TestingModeIsOn bool
var yam *config.YamlConfig
var proxyURL *url.URL
var shareHTTPClient *http.Client

type MenuEntry struct {
	Title   string
	Handler func()
}

var menus = []MenuEntry{
	{Title: "(POST) > 🔒 	| Basic Auth Fuzzing - URLs Only", Handler: fuzzBasicAuthPost},
	{Title: "(GET)  > 🔒 	| Basic Auth Fuzzing - URLs Only", Handler: fuzzBasicAuthGet},
	{Title: "(GET)  > 🔓	| Dir Bust - URLs + Paths", Handler: DirBustURLAndPaths},
	{Title: "(POST) > 🔓🔒 | GraphQl fuzzing - URLs + Paths", Handler: fuzzGraphQL},
	{Title: "Infra  > 🖥️ 	 | Host HTTPS Server at 443", Handler: hostHTTPServer},
}

func main() {

	IgniteTheEngine()
	runCLIMenu()
}

// This should be called first. This initialize all neccessary variables and configuration
func IgniteTheEngine() {
	loadConfigFromFiles()
	utils.SetTestingMode(TestingModeIsOn)
	shareHTTPClient = utils.InitHTTPClient(proxyURL)
}

func runCLIMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(utils.GreenText(title))
		for i, e := range menus {
			fmt.Printf("%d) %s \n", i+1, e.Title)
		}

		fmt.Print(utils.GreenText("\nPick Your choice >> "))

		if !scanner.Scan() {
			break
		}

		fmt.Println("")

		choice := strings.TrimSpace(scanner.Text())
		index, error := strconv.Atoi(choice)
		if error != nil || index < 1 || index > len(menus) {
			fmt.Println("Invalid choice")
			continue
		}

		menus[index-1].Handler()
	}
}

func loadConfigFromFiles() {
	// TestingModeIsOn flag bootstrapped. To enable the testing mode, invoke -> go go-nut.go --test
	flag.BoolVar(&TestingModeIsOn, "test", false, "enable testing mode")
	flag.Parse()

	//loading configuration from yaml file
	yam_f_name := config.YAMLFileName
	if TestingModeIsOn {
		// to print Debug function output, we need to set DebugOn to true. This is a bit hacky but it works for now. Need to refactor in the future.
		utils.DebugOn = true
		utils.Debug("The testing mode is on! Loading test-config.yaml...")
		yam_f_name = config.YAMLFileNameForTesting
	}

	// somehow you have to load local first then assign to the global yam to be reusable within the main package
	loaded, err := LoadYamlConfig(yam_f_name)
	if err != nil {
		utils.Error("[x] Unable to load yaml config file!")
	}
	yam = loaded

	url, err := url.Parse(yam.ProxyURLString)
	if err != nil {
		utils.Error("Failed to load the proxy URL!")
	}
	proxyURL = url

	utils.DebugOn = yam.Debug
}

func invokeRequestWithBasicAuth(client *http.Client, url string, method config.HTTPMethod, c config.ContentType, body io.Reader, userCredential *config.UserCredential, outputMode config.PrintOutputMode) {

	req, err := http.NewRequest(method.String(), url, body)
	req.Header.Set("Content-Type", c.String())

	if userCredential != nil {
		req.SetBasicAuth(userCredential.Username, userCredential.Password)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(utils.RedText("[x] HTTP Client error!\n"), err)
	}

	fmt.Println(utils.BlueText("\n[i] Request Sent:"), req.URL)

	defer resp.Body.Close()

	switch outputMode {
	case config.PrintRespHeaders:
		utils.PrintHeaders(resp)
	case config.PrintRespBody:
		utils.PrintBody(resp)
	}
}

func invokeHTTPGET(client *http.Client, url string, outputMode config.PrintOutputMode) {

	req, err := http.NewRequest(config.HTTPMethodGet.String(), url, nil)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(utils.RedText("[x] HTTP Client error!\n"), err)
	}

	msg := fmt.Sprintf("Request Sent: %s (status %d)", req.URL, resp.StatusCode)
	utils.Info(msg)

	defer resp.Body.Close()

	switch outputMode {
	case config.PrintRespHeaders:
		utils.PrintHeaders(resp)
	case config.PrintRespBody:
		utils.PrintBody(resp)
	default:

	}
}

func invokeHTTPPOST(client *http.Client, url string, c config.ContentType, body io.Reader, outputMode config.PrintOutputMode, cookies string) {

	req, err := http.NewRequest(config.HTTPMethodPost.String(), url, body)
	req.Header.Set("Content-Type", c.String())
	if cookies != "" {
		req.Header.Set("Cookie", cookies)
	} else {
		utils.Error("Cookies file " + utils.GetCookiesFileName() + " is empty!. This will be unauthenticated request without cookies.")
	}

	if err != nil {
		fmt.Println(utils.RedText("[x] HTTP Client error!\n"), err)
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(utils.RedText("[x] HTTP Client error!\n"), err)
		return
	}

	defer resp.Body.Close()

	msg := fmt.Sprintf("Request Sent: %s (status %d)", req.URL, resp.StatusCode)
	utils.Info(msg)

	switch outputMode {
	case config.PrintRespHeaders:
		utils.PrintHeaders(resp)
	case config.PrintRespBody:
		utils.PrintBody(resp)
	default:

	}
}

// loading config.yaml file from config/ folder
func LoadYamlConfig(path string) (*config.YamlConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var yamyam config.YamlConfig
	if err := yaml.Unmarshal(data, &yamyam); err != nil {
		return nil, err
	}

	if TestingModeIsOn || yamyam.Debug {
		c, _ := json.MarshalIndent(yamyam, "", "  ")
		utils.Debug("YAML Config loaded: " + string(c))
	}

	return &yamyam, nil
}

// TODO all hard coded. same code from POST method
func fuzzBasicAuthGet() {

	jobs, wg := utils.RaiseMinions(100)
	// getting urls from urls.txt file loaded by commonFileLoader function
	for fr := range utils.LoadFileStream(utils.GetUrlFileName()) {
		// fmt.Printf("#%d %s\n", fr.Row, fr.Line)
		jobs <- func() {
			invokeRequestWithBasicAuth(shareHTTPClient, fr.Text, config.HTTPMethodGet, config.ContentTypeXml, strings.NewReader(yam.BodyXml), &yam.ValidCred, config.PrintRespBody)
		}
		jobs <- func() {
			invokeRequestWithBasicAuth(shareHTTPClient, fr.Text, config.HTTPMethodGet, config.ContentTypeXml, strings.NewReader(yam.BodyXml), &yam.InvalidCred, config.PrintRespHeaders)
		}
		jobs <- func() {
			invokeRequestWithBasicAuth(shareHTTPClient, fr.Text, config.HTTPMethodGet, config.ContentTypeXml, nil, nil, config.PrintRespWhole)
		}

	}
	close(jobs)
	wg.Wait()
}

func fuzzBasicAuthPost() {
	// getting urls from urls.txt file loaded by commonFileLoader function
	jobs, wg := utils.RaiseMinions(100)

	for fr := range utils.LoadFileStream(utils.GetUrlFileName()) {
		// fmt.Printf("#%d %s\n", fr.Row, fr.Line)
		jobs <- func() {
			invokeRequestWithBasicAuth(shareHTTPClient, fr.Text, config.HTTPMethodPost, config.ContentTypeXml, strings.NewReader(yam.BodyXml), &yam.ValidCred, config.PrintRespBody)
		}
		jobs <- func() {
			invokeRequestWithBasicAuth(shareHTTPClient, fr.Text, config.HTTPMethodPost, config.ContentTypeXml, strings.NewReader(yam.BodyXml), &yam.InvalidCred, config.PrintRespHeaders)
		}
		jobs <- func() {
			invokeRequestWithBasicAuth(shareHTTPClient, fr.Text, config.HTTPMethodPost, config.ContentTypeXml, nil, nil, config.PrintRespWhole)
		}
	}

	close(jobs)
	wg.Wait()
}

func DirBustURLAndPaths() {
	utils.Processing("Loading URLs and Paths from files...")
	urlFr := utils.LoadFileClassically(utils.GetUrlFileName())
	pathFr := utils.LoadFileClassically(utils.GetPathFileName())
	var targetUrls []config.FileRow

	utils.Processing("Generating target URLs by combining URLs and paths...")
	row := 0
	for _, pFr := range pathFr {
		for _, ur := range urlFr {
			row++
			targetUrl := strings.TrimRight(ur.Text, "/") + "/" + strings.TrimLeft(pFr.Text, "/")
			targetUrls = append(targetUrls, config.FileRow{
				Row:  row,
				Text: targetUrl,
			})
		}
	}

	jobs, wg := utils.RaiseMinions(100)

	utils.Processing("Sending requests to the server(s)...\n")
	for _, cFr := range targetUrls {
		// need shadow variable to avoid closure issue
		url := cFr.Text
		jobs <- func() {
			invokeHTTPGET(shareHTTPClient, url, config.PrintNone)
		}
	}

	close(jobs)
	wg.Wait()
	utils.Success(fmt.Sprintf("Total request: %d", row))
	utils.Success("Done")
}

func promptInputs(prompts []string, next func(map[string]string)) {
	inputs := make(map[string]string)
	for _, p := range prompts {
		fmt.Print(utils.GreenText(p + ">>> "))
		var value string
		fmt.Scanln(&value)
		inputs[p] = value
	}

	next(inputs)
}

func promptURLAndPathFiles() {
	promptInputs([]string{"URL File Path >>", "Paths File Path >>"}, func(m map[string]string) {
		fmt.Println("URL File:", m["URL File Path >>"])
		fmt.Println("Paths File:", m["Paths File Path >>"])
	})
}

func fuzzGraphQL() {

	utils.Processing("Loading cookies from file...")
	cookies := utils.LoadWholeFileAsString(utils.GetCookiesFileName())

	utils.Processing("Loading URLs and Paths from files...")
	urlFr := utils.LoadFileClassically(utils.GetUrlFileName())
	pathFr := utils.LoadFileClassically(utils.GetPathFileName())
	var targetUrls []config.FileRow

	utils.Processing("Generating target URLs by combining URLs and paths...")
	row := 0
	for _, pFr := range pathFr {
		for _, ur := range urlFr {
			row++
			targetUrl := strings.TrimRight(ur.Text, "/") + "/" + strings.TrimLeft(pFr.Text, "/")
			targetUrls = append(targetUrls, config.FileRow{
				Row:  row,
				Text: targetUrl,
			})
		}
	}

	data := attack.GraphqlQueries

	jobs, wg := utils.RaiseMinions(100)

	utils.Processing("Sending requests to the server(s)...\n")
	total := 0
	for _, cFr := range targetUrls {
		// need shadow variable to avoid closure issue
		url := cFr.Text
		// preparing the body data for POST request
		for _, v := range data {
			jsonBytes, _ := json.Marshal(v)
			reader := bytes.NewBuffer(jsonBytes)
			jobs <- func() {
				total++
				invokeHTTPPOST(shareHTTPClient, url, config.ContentTypeJson, reader, config.PrintRespBody, cookies)
			}
		}
	}

	close(jobs)
	wg.Wait()
	utils.Success(fmt.Sprintf("Total request: %d", total))
	utils.Success("Done")

}

func hostHTTPServer() {
	utils.Processing("Starting HTTP server ...")
	utils.Debug("HTTPServerPort configuredin yaml file is " + strconv.Itoa(yam.HTTPServerPort))
	utils.ServeHTTPServer(yam.HTTPServerPort)
}
