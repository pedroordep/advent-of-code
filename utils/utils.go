package utils

import (
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

//go:embed session.txt
var sessionCookie string

func GetInputFile(year, day int) string {
	_, file, _, _ := runtime.Caller(0)
	runtimePath := filepath.Dir(file)

	inputFilePath := fmt.Sprintf("%s/../%d/%d/input.txt", runtimePath, year, day)

	_, err := os.Stat(inputFilePath)
	if err == nil {
		fmt.Println("Input file found locally:", inputFilePath)
		body, _ := os.ReadFile(inputFilePath)
		return string(body)
	}

	path := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, _ := http.NewRequest("GET", path, nil)

	req.AddCookie(&http.Cookie{Name: "session", Value: sessionCookie})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic("Could not fetch input file")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic("Could not read response body")
		}

		os.WriteFile(inputFilePath, bodyBytes, 0644)
		fmt.Println("Saved file locally:", inputFilePath)

		return string(bodyBytes)
	}

	fmt.Println("Could not fetch input file, status code: " + resp.Status)

	return ""
}
