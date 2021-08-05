package solution

import (
	"math/rand"
	"time"
)

// simple take the array and shuffle it
// right now learnring how to use "flags" so will
// be writing very inefficient code for randomization
func GetRandomizedQuiz(quizLines [][]string) [][]string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(quizLines); n > 0; n-- {
		randIndex := r.Intn(len(quizLines))
		// well this swap is amazingly done by go. lol
		quizLines[n-1], quizLines[randIndex] = quizLines[randIndex], quizLines[n-1]
	}
	return quizLines
}
