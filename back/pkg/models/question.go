package models

// Question represents a question
type Question struct {
	ID            int    `json:"id"`
	Text          string `json:"text"`
	CorrectAnswer string `json:"correctAnswer"`
	Level         int    `json:"level"`
}

// AnswerRequest represents the request to check an answer
type AnswerRequest struct {
	QuestionID int    `json:"question_id"`
	Answer     string `json:"answer"`
}

// ScoreResponse represents the response of the score
type ScoreResponse struct {
	Score int `json:"score"`
}
