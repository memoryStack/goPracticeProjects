package solution

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func timer(duration int, timerChan chan int) {
	// simply sleep for some seconds and then return
	time.Sleep(time.Duration(duration) * time.Second)
	timerChan <- 1
}

func PartB(time int) {
	// Interesting Note: the file problems.csv should be present in the directry from which we are going to run
	// "main.go" file to run this program
	// Rule: i will always run program from the root directry (where README.md) is defiend
	f, err := os.Open("problems.csv")
	if err != nil {
		return
	}
	defer f.Close()

	lines, _ := csv.NewReader(f).ReadAll()
	totalCorrectAnsweres := 0

	var temp string
	fmt.Printf("Press Enter to start the quiz:\n")
	fmt.Scanf("%s", &temp)

	timerChan := make(chan int)
	quizChan := make(chan int)

	go timer(time, timerChan)

	go func() {
		for index, value := range lines {
			fmt.Printf("Problem #%d:\n", index+1)
			fmt.Printf("%s = ", value[0])
			var ans int
			fmt.Scanf("%d", &ans)
			if intVar, _ := strconv.Atoi(value[1]); ans == intVar {
				totalCorrectAnsweres++
			}
		}
		quizChan <- 1
	}()

	for {
		select {
		case <-timerChan:
			fmt.Printf("timer is over, ending your test\n")
			break
		case <-quizChan:
			fmt.Printf("attempted the test before time, congratulations\n")
			break
		}
		fmt.Printf("score: %d / %d\n", totalCorrectAnsweres, len(lines))
		return
	}

}
