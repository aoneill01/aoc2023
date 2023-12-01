package util

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func ReadDay(day int) []string {
	godotenv.Load()

	filename := cacheFilename(day)
	var s string

	if fileExists(filename) {
		input, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		s = string(input)
	} else {
		s = downloadInput(day)
		cacheInput(s, day)
	}

	ss := strings.Split(s, "\n")
	return ss[:len(ss)-1]
}

func downloadInput(day int) string {
	url := fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day)
	c := http.Client{Timeout: time.Duration(3) * time.Second}

	session := &http.Cookie{
		Name:   "session",
		Value:  os.Getenv("AOC_SESSION"),
		MaxAge: 0,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error with HTTP request: %s", err.Error())
	}
	req.AddCookie(session)

	resp, err := c.Do(req)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("Warning: Could not authenticate with AOC server")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	return string(body)
}

func cacheFilename(day int) string {
	return fmt.Sprintf("./inputs/day-%02d.txt", day)
}

func cacheInput(contents string, day int) {
	err := os.WriteFile(cacheFilename(day), []byte(contents), 0644)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
