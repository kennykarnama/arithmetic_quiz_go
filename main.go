package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kennykarnama/arithmetic_quiz_go/quiz"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("Missing path file")
	}

	quiz, err := quiz.NewQuizFromCsvFile("sample_question.csv")
	if err != nil {
		log.Fatal(err)
	}
	idx := 0
	correctAnswer := 0
	totalQuestion := len(quiz.Questions)
	for idx < len(quiz.Questions) {
		q := quiz.Pick(idx)
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("%s\n", q.QuestionDescription)
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\n")
		text = strings.TrimSpace(text)
		keyAnswer := strings.TrimRight(q.KeyAnswer, "\n")
		if text == keyAnswer {
			correctAnswer++
		} else {
			fmt.Printf("%s %s \n", text, keyAnswer)
		}
		idx++
	}
	fmt.Printf("****** RESULT *****\n")
	fmt.Printf("Correct Answer %d\n", correctAnswer)
	fmt.Printf("Total questions %d\n", totalQuestion)
	fmt.Printf("****** END RESULT *****\n")
}
