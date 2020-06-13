package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type pageData struct {
	Content string
	Title   string
}

func main() {

	startTime := time.Now()

	fileFlag := flag.String("file", "first-post.txt", "used to define input text")
	dirFlag := flag.String("dir", "none", "if spedified, generates for all .txt in the directory")
	outDirFlag := flag.String("out", "output_html/", "SSG output directory")
	flag.Parse()

	var fileCount int = 0
	var fileSize float64 = 0

	if *dirFlag == "none" {
		runForFile(*fileFlag, "text_directory/", &fileCount, &fileSize)
	} else {
		var directory string = *dirFlag
		runForDir(directory, *outDirFlag, &fileCount, &fileSize)
	}

	elapsedTimeFloat := time.Since(startTime).Seconds()
	elapsedTime := strconv.FormatFloat(elapsedTimeFloat, 'f', 6, 64)
	elapsedTime = elapsedTime[0 : strings.Index(elapsedTime, ".")+3]

	sizeString := strconv.FormatFloat(fileSize, 'f', -1, 32)
	sizeString = sizeString[0 : strings.Index(sizeString, ".")+3]

	fmt.Printf(
		"\n\033[32;1mSuccess!\033[0m Created \033[1m%d\033[0m files with a size of %skb in %s seconds!\n",
		fileCount,
		sizeString,
		elapsedTime,
	)

}

func runForDir(directory, out string, fileCountPtr *int, fileSizePtr *float64) {
	if directory[len(directory)-1] != "/"[0] {
		directory += "/"
	}

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() == false {
			runForFile(file.Name(), directory, fileCountPtr, fileSizePtr)
		} else {
			runForDir(directory+"/"+file.Name(), out, fileCountPtr, fileSizePtr)
		}
	}
}

func runForFile(fileFlag, directory string, fileCountPtr *int, fileSizePtr *float64) {
	var fileName string = fileFlag
	if fileName[strings.Index(fileFlag, "."):len(fileFlag)] != ".txt" {
		return
	}

	fileName = fileName[0:strings.Index(fileFlag, ".")] + ".html"

	var textData string = readFile(directory + fileFlag)
	renderTemplate("template.tmpl", textData, fileName, fileCountPtr, fileSizePtr)
}

func readFile(inPath string) string {
	fileContents, err := ioutil.ReadFile(inPath)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func renderTemplate(tPath, textData, fileName string, fileCountPtr *int, fileSizePtr *float64) {
	paths := []string{
		tPath,
	}

	f, err := os.Create("output_html/" + fileName)
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

	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}
	*fileCountPtr++
	*fileSizePtr += float64(fi.Size()) / float64(1000) //1000 or 1024? I think 1024 is kibabite but honestly I don't really care

	f.Close()

	// bytesToWrite := []byte(*t)
	// err = ioutil.WriteFile("whatever.html", bytesToWrite, 0644) // the hell is 0644??
	// if err != nil {
	// 	panic(err)
	// }
}
