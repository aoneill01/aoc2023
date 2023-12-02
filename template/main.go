package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"text/template"
)

//go:embed main.go.tmpl
var t1 string

//go:embed main_test.go.tmpl
var t2 string

type Data struct {
	Day int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing day argument")
		return
	}

	day, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Day must be a number")
		return
	}

	d := Data{day}

	dir := fmt.Sprintf("day%02d", d.Day)

	if err := os.Mkdir(dir, 0755); err != nil {
		fmt.Println("Could not create the directory")
		return
	}

	writeTemplate(t1, fmt.Sprintf("%s/main.go", dir), d)
	writeTemplate(t2, fmt.Sprintf("%s/main_test.go", dir), d)
}

func writeTemplate(t, filename string, d Data) {
	mainTemplate := template.New("main")
	mainTemplate.Parse(t)

	if f, err := os.Create(filename); err == nil {
		defer f.Close()
		mainTemplate.Execute(f, d)
	} else {
		fmt.Println("Error creating file")
		return
	}
}
