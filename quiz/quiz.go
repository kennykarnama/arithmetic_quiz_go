package quiz

import (
	"github.com/kennykarnama/arithmetic_quiz_go/model/simple_question"
	"github.com/kennykarnama/arithmetic_quiz_go/parserutil"
)

type Quiz struct {
	Questions []simple_question.QuestionItem
}

// Construct new quiz from csv file
func NewQuizFromCsvFile(fileloc string) (*Quiz, error) {
	parser, err := parserutil.NewCsvFile(fileloc)

	if err != nil {
		return nil, err
	}

	records, err := parser.QueryRecords()

	if err != nil {
		return nil, err
	}

	var questions []simple_question.QuestionItem
	// Csv File must omit header
	for _, record := range records {
		q := simple_question.QuestionItem{
			QuestionDescription: record[0],
			KeyAnswer:           record[1],
		}
		questions = append(questions, q)
	}

	newQuiz := Quiz{
		Questions: questions,
	}

	return &newQuiz, nil
}

// Pick question at certain index
// Input parameter idx int
// Return QuestionItem
func (quiz *Quiz) Pick(idx int) simple_question.QuestionItem {
	return quiz.Questions[idx]
}
