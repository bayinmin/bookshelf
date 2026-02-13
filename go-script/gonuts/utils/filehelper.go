package utils

import (
	"bufio"
	"gonuts/config"
	"os"
	"strings"
)

// load file and read concurrently
func LoadFileStream(filename string) chan config.FileRow {

	ch := make(chan config.FileRow)

	go func() {
		defer close(ch)

		file, err := os.Open(filename)
		if err != nil {
			Error("Error loading the file in stream/channel mode!")
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		row := 0
		for scanner.Scan() {
			row++
			fr := config.FileRow{
				Row:  row,
				Text: scanner.Text(),
			}
			ch <- fr
		}
	}()

	return ch
}

func LoadFileClassically(fname string) []config.FileRow {
	var out []config.FileRow

	file, err := os.Open(fname)
	if err != nil {
		Error("Error loading the file classically!")
		return out
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		row++
		out = append(out, config.FileRow{
			Row:  row,
			Text: scanner.Text(),
		})
	}
	return out

}

func LoadWholeFileAsString(fname string) string {
	file, err := os.ReadFile(fname)
	if err != nil {
		Error("Error loading LoadWholeFileAsString!")
		return ""
	}
	return strings.TrimSpace(string(file))
}
