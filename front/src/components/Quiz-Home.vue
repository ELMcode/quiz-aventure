<template>
  <div>
    <h1>Quiz Aventure</h1>

    <h2>Bienvenue dans Quiz Aventure !</h2>

    <p>Quizz de culture générale</p>
    <p>Réponse correcte = 10 points !</p>
    <p>Réponse incorrecte = - 5 points !</p>

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
      currentQuestionIndex: 0,
      userAnswer: ""
    };
  },
  computed: {
    currentQuestion() {
      return this.questions[this.currentQuestionIndex];
    }
  },
  methods: {
    fetchQuestions() {
      fetch('http://localhost:8081/questions')
          .then(response => response.json())
          .then(data => {
            this.questions = data;
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
            if (this.currentQuestionIndex >= this.questions.length) {
              alert("Quiz terminé! Votre score est: " + this.score);
              this.currentQuestionIndex = 0
              this.score = 0;
            }
          });
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
