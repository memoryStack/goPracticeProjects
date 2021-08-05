package main

import (
	"flag"
	"strconv"

	solution "github.com/goPracticeProjects/Problem_1/solution"
)

func main() {

	// Note that the flag package requires all flags to appear before positional
	// arguments (otherwise the flags will be interpreted as positional arguments)
	shufflePtr := flag.Bool("shuffle", false, "a bool")
	flag.Parse()

	args := flag.Args() // use args like this while using flags
	if args[0] == "A" {
		solution.PartA(*shufflePtr)
	} else {
		// timer will be taken from commandline-arguments
		time, err := strconv.Atoi(args[1])
		if err != nil {
			return
		}
		solution.PartB(time, *shufflePtr)
	}
}
