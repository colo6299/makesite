package main

import (
	"io/ioutil"
	"os"
	"text/template"
)

type PageData struct {
	Content string
}

func main() {
	var textData string = readFile()
	renderTemplate("template.tmpl", textData)
}

func readFile() string {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func renderTemplate(tPath, textData string) {
	paths := []string{
		tPath,
	}

	f, err := os.Create("first-post.html")
	if err != nil {
		panic(err)
	}

	t, err := template.New(tPath).ParseFiles(paths...)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, PageData{textData})
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
