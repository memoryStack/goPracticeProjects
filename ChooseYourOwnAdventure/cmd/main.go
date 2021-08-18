package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/goPracticeProjects/ChooseYourOwnAdventure/solution"
)

func main() {
	// receive the port on which server would be starting
	port := flag.Int("port", 8080, "port on which server would start")
	// receive the file name in flags from commandline which will be parsed for our story online
	fileName := flag.String("file", "jsonFile.json", "json file with the CYOA story")

	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		// author says that panic is not a great idea to handle the errors in go but for now let's use
		// it anyway. Ques -> why panics are not a nice way to handle errors in go ?
		panic(err)
	}

	// file is opened and now let's parse the json and store in a variable
	// Note: in order for the auto-import to work the folder name and the package name
	// should be same else auto-import just wouldn't work
	story, err := solution.JSONStory(file)
	if err != nil {
		panic(err)
	}

	handler := solution.NewHandler(story)
	fmt.Printf("Server would be starting on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handler))

	// using "+v" option to print out all the keys and the contents to get
	// a better clearity of the data structure
	// fmt.Printf("%+v\n", story)
}
