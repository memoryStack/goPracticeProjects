package main

import (
	"os"
	"strconv"

	solution "github.com/goPracticeProjects/Problem_1/solution"
)

func main() {

	args := os.Args
	if args[1] == "A" {
		solution.PartA()
	} else {
		// timer will be taken from commandline-arguments
		time, err := strconv.Atoi(args[2])
		if err != nil {
			return
		}
		solution.PartB(time)
	}
}
