package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/goPracticeProjects/HTMLLinkParser/solution"
	"golang.org/x/net/html"
)

func main() {
	fileNum := flag.Int("fileNum", 1, "which file number to parse [1-4]")
	flag.Parse()
	fileName := fmt.Sprintf("htmlFiles/html%d.html", *fileNum)
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	// TODO: find out that how is this "file" is io.Reader type ?
	// are these types similar ??
	root, err := html.Parse(file)
	if err != nil {
		panic(err)
	}

	links := solution.GetLinks(root)
	fmt.Printf("%v\n", links)
}
