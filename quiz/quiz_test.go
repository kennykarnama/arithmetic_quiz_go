package quiz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	quiz, err := NewQuizFromCsvFile("sample_question.csv")
	assert.Nil(t, err)
	listQuestions := quiz.Questions
	assert.NotNil(t, listQuestions)
	assert.Greater(t, len(listQuestions), 0)
}

func TestLoadFileNotFound(t *testing.T) {
	_, err := NewQuizFromCsvFile("invalid_field_length.csv")
	assert.NotNil(t, err)
}

func TestLoadFileFoundButEmpty(t *testing.T) {
	_, err := NewQuizFromCsvFile("empty_question.csv")
	assert.NotNil(t, err)
}

func TestPickAfterLoad(t *testing.T) {
	quiz, err := NewQuizFromCsvFile("sample_question.csv")
	assert.Nil(t, err)
	firstQuestion := quiz.Pick(0)
	assert.NotNil(t, firstQuestion)
	assert.NotEmpty(t, firstQuestion.QuestionDescription)
	assert.NotEmpty(t, firstQuestion.KeyAnswer)
}
