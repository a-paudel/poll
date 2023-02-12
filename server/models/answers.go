package models

type Answer struct {
	ID     uint   `json:"id"`
	Answer string `json:"answer"`
	Votes  int    `json:"votes"`

	QuestionId uint     `json:"questionId"`
	Question   Question `json:"-"`
}
