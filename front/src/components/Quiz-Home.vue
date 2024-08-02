<template>
  <div>
    <h1>Quiz Aventure</h1>

    <h2 v-if="currentLevelName">Niveau: {{ currentLevelName }}</h2>

    <div v-if="currentQuestion">
      <h2>{{ currentQuestion.text }}</h2>
      <input v-model="userAnswer" placeholder="Votre réponse" />
      <button @click="checkAnswer">Valider</button>
    </div>
    <p>Score: {{ score }}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      score: 0,
      questions: [],
      levels: ["Débutant", "Intermédiaire", "Expert"],
      currentLevel: 1,
      currentQuestionIndex: 0,
      userAnswer: ""
    };
  },
  computed: {
    currentQuestion() {
      return this.questions.filter(q => q.level === this.currentLevel)[this.currentQuestionIndex];
    },
    currentLevelName() {
      return this.levels[this.currentLevel - 1];
    }
  },
  methods: {
    fetchQuestions() {
      fetch('http://localhost:8081/questions')
          .then(response => response.json())
          .then(data => {
            this.questions = data;
            this.restoreProgress();
          });
    },
    checkAnswer() {
      fetch('http://localhost:8081/check-answer', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          question_id: this.currentQuestion.id,
          answer: this.userAnswer
        })
      })
          .then(response => response.json())
          .then(data => {
            this.score = data.score;
            this.userAnswer = "";
            this.currentQuestionIndex++;
            const currentLevelQuestions = this.questions.filter(q => q.level === this.currentLevel);
            if (this.currentQuestionIndex >= currentLevelQuestions.length) {
              this.currentQuestionIndex = 0;
              this.currentLevel++;
              if (this.currentLevel > this.levels.length) {
                alert("Quiz terminé! Votre score final est: " + this.score);
                this.resetQuiz();
              }
            }
            this.saveProgress();
          });
    },
    resetQuiz() {
      fetch('http://localhost:8081/reset-score', {
        method: 'POST'
      })
          .then(() => {
            this.currentLevel = 1;
            this.currentQuestionIndex = 0;
            this.score = 0;
            this.saveProgress();
          })
          .catch(error => {
            console.error('Error resetting score:', error);
          });
    },
    saveProgress() {
      localStorage.setItem('quizCurrentLevel', this.currentLevel);
      localStorage.setItem('quizCurrentQuestionIndex', this.currentQuestionIndex);
      localStorage.setItem('quizScore', this.score);
    },
    restoreProgress() {
      const savedLevel = localStorage.getItem('quizCurrentLevel');
      const savedQuestionIndex = localStorage.getItem('quizCurrentQuestionIndex');
      const savedScore = localStorage.getItem('quizScore');

      if (savedLevel !== null) {
        this.currentLevel = parseInt(savedLevel, 10);
      }
      if (savedQuestionIndex !== null) {
        this.currentQuestionIndex = parseInt(savedQuestionIndex, 10);
      }
      if (savedScore !== null) {
        this.score = parseInt(savedScore, 10);
      }
    }
  },
  mounted() {
    this.fetchQuestions();
  }
};
</script>

<style scoped>
h1 {
  color: #2c3e50;
}
</style>
