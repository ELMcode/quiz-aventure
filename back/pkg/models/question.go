package models

type Question struct {
	ID            int    `json:"id"`
	Text          string `json:"text"`
	CorrectAnswer string `json:"correctAnswer"`
}
