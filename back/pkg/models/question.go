package models

type Question struct {
	ID            int    `json:"id"`
	Text          string `json:"text"`
	CorrectAnswer string `json:"correctAnswer"`
	Level         int    `json:"level"`
}

type AnswerRequest struct {
	QuestionID int    `json:"question_id"`
	Answer     string `json:"answer"`
}

type ScoreResponse struct {
	Score int `json:"score"`
}
