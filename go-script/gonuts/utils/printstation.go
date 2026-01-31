package utils

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

var DebugOn bool

func Debug(s string) {
	if DebugOn {
		PrintNormal(s)
	}
}

func Success(s string) {
	fmt.Println(GreenText("[✓]") + " " + s)
}

func Error(s string) {
	fmt.Println(GreenText("[X]") + " " + s)
}

func PrintNormal(s string) {
	fmt.Println(s)
}

func Processing(s string) {
	fmt.Println(PurpleText("[*]") + " " + s)
}

func Info(s string) {
	fmt.Println(BlueText("[i]") + " " + s)
}

func RedText(s string) string {
	return "\033[31m" + s + "\033[0m"
}

func GreenText(s string) string {
	return "\033[32m" + s + "\033[0m"
}

func BlueText(s string) string {
	return "\033[34m" + s + "\033[0m"
}

func PrintHeaders(resp *http.Response) {
	fmt.Println("")
	fmt.Println("Status:", resp.Status)
	for k, v := range resp.Header {
		fmt.Printf("%s: %s\n", k, strings.Join(v, ", "))
	}
}

func PurpleText(s string) string {
	return "\033[35m" + s + "\033[0m"
}

func PrintBody(resp *http.Response) {
	fmt.Println("")
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
