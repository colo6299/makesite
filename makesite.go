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
	Title   string
}

func main() {

	fileFlag := flag.String("file", "first-post.txt", "used to define input text")
	dirFlag := flag.String("dir", "none", "if spedified, generates for all .txt in the directory")
	flag.Parse()
	if *dirFlag == "none" {
		runForFile(*fileFlag, "text_directory/")
	} else {
		var directory string = *dirFlag
		if directory[len(directory)-1] != "/"[0] {
			directory += "/"
		}

		files, err := ioutil.ReadDir(directory)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			if file.IsDir() == false {
				runForFile(file.Name(), directory)
			}
		}
	}
}

func runForFile(fileFlag, directory string) {
	var fileName string = fileFlag
	fileName = fileName[0:strings.Index(fileFlag, ".")] + ".html"

	var textData string = readFile(directory + fileFlag)
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

	rawName := fileName[0:strings.Index(fileName, ".")]

	err = t.Execute(f, pageData{textData, rawName})
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
