package simple_question

const (
	Add      = "+"
	Minus    = "-"
	Multiply = "*"
	Divide   = "/"
)

type AnsweredItem struct {
	Answer int
}

type QuestionItem struct {
	QuestionDescription string
	KeyAnswer           string
}
