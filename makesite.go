package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type pageData struct {
	Content string
}

func main() {

	fileFlag := flag.String("file", "first-post.txt", "used to define input text")
	flag.Parse()
	var fileName string = *fileFlag
	fileName = fileName[0:strings.Index(*fileFlag, ".")] + ".html"

	var textData string = readFile(*fileFlag)
	renderTemplate("template.tmpl", textData, fileName)
}

func readFile(inPath string) string {
	fileContents, err := ioutil.ReadFile(inPath)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func renderTemplate(tPath, textData, fileName string) {
	paths := []string{
		tPath,
	}

	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	t, err := template.New(tPath).ParseFiles(paths...)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, pageData{textData})
	if err != nil {
		panic(err)
	}

	f.Close()

	// bytesToWrite := []byte(*t)
	// err = ioutil.WriteFile("whatever.html", bytesToWrite, 0644) // the hell is 0644??
	// if err != nil {
	// 	panic(err)
	// }
}
