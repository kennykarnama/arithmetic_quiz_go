package main

import (
	"fmt"
	"log"

	"github.com/kennykarnama/arithmetic_quiz_go/quiz"
)

func main() {
	quiz, err := quiz.NewQuizFromCsvFile("sample_question.csv")
	if err != nil {
		log.Fatal(err)
	}
	idx := 0
	for idx < len(quiz.Questions) {
		q := quiz.Pick(idx)
		fmt.Printf("%s\n", q.QuestionDescription)
		idx++
	}
}
