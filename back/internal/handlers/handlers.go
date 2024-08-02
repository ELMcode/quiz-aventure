package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"quiz-aventure/internal/quiz"
	"quiz-aventure/pkg/models"
)

var (
	questions []models.Question
	score     int
	mutex     sync.Mutex
)

func init() {
	questions = quiz.LoadQuestions()
	score = 0
}

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func CheckAnswer(w http.ResponseWriter, r *http.Request) {
	var req models.AnswerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	correct := false
	for _, question := range questions {
		if question.ID == req.QuestionID {
			if question.CorrectAnswer == req.Answer {
				correct = true
				break
			}
		}
	}

	mutex.Lock()
	if correct {
		score += 10
	} else {
		score -= 5
	}
	currentScore := score
	mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		models.ScoreResponse{
			Score: currentScore})
}

func ResetScore(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	score = 0
	mutex.Unlock()
	w.WriteHeader(http.StatusOK)
}
