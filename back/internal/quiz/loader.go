package quiz

import (
	"encoding/json"
	"log"
	"os"

	"quiz-aventure/pkg/models"
)

// LoadQuestions loads the questions from the JSON file
func LoadQuestions() []models.Question {
	file, err := os.ReadFile("internal/quiz/questions.json")
	if err != nil {
		log.Fatalf("could not read questions file: %v\n", err)
	}

	var questions []models.Question
	if err := json.Unmarshal(file, &questions); err != nil {
		log.Fatalf("could not unmarshal questions: %v\n", err)
	}

	return questions
}
