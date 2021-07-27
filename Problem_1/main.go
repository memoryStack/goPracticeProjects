// package main

// import (
//    "fmt"
//    "encoding/csv"
// )

// func main() {

// }

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("problems.csv")
	if err != nil {
		return
	}
	defer f.Close()

	lines, _ := csv.NewReader(f).ReadAll()
	totalCorrectAnsweres := 0

	for index, value := range lines {
		fmt.Printf("Problem #%d:\n", index+1)
		fmt.Printf("%s = ", value[0])
		var ans int
		fmt.Scanf("%d", &ans)
		if intVar, _ := strconv.Atoi(value[1]); ans == intVar {
			totalCorrectAnsweres++
		}
	}

	fmt.Printf("score: %d / %d\n", totalCorrectAnsweres, len(lines))
}

/*
	About "defer"
		A defer statement defers the execution of a function until the surrounding function returns.
		The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
*/
