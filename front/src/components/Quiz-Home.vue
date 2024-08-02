<template>
  <div>
    <div v-if="!pseudo">
      <PseudoForm @submit="setPseudo" />
    </div>
    <div v-else>
      <h1 class="animate__animated animate__fadeIn">Quiz Aventure</h1>

      <h2 v-if="currentLevelName" class="animate__animated animate__fadeIn">Niveau: {{ currentLevelName }}</h2>

      <transition name="fade" mode="out-in">
        <div v-if="currentQuestion" :key="currentQuestion.id" class="animate__animated animate__fadeIn">
          <h2>{{ currentQuestion.text }}</h2>
          <input v-model="userAnswer" placeholder="Votre réponse" />
          <button @click="checkAnswer">Valider</button>
          <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
        </div>
      </transition>
      <p>Pseudo: {{ pseudo }} | Score: {{ score }}</p>
      <button @click="resetQuiz">Recommencer</button>
      <button @click="changePlayer">Changer de joueur</button>
    </div>
  </div>
</template>

<script>
import PseudoForm from './PseudoForm.vue';
import confetti from 'canvas-confetti';

export default {
  components: {
    PseudoForm
  },
  data() {
    return {
      pseudo: '',
      score: 0,
      questions: [],
      levels: ["Débutant", "Intermédiaire", "Expert"],
      currentLevel: 1,
      currentQuestionIndex: 0,
      userAnswer: "",
      errorMessage: null
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
    setPseudo(pseudo) {
      this.pseudo = pseudo;
      this.saveProgress();
    },
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
            const previousScore = this.score;
            this.score = data.score;
            if (this.score > previousScore) {
              this.launchConfetti();
              this.errorMessage = null;
            } else {
              this.errorMessage = 'Mauvaise réponse !';
              setTimeout(() => {
                this.errorMessage = null;
              }, 3000);

            }
            this.userAnswer = "";
            this.currentQuestionIndex++;
            const currentLevelQuestions = this.questions.filter(q => q.level === this.currentLevel);
            if (this.currentQuestionIndex >= currentLevelQuestions.length) {
              this.currentQuestionIndex = 0;
              this.currentLevel++;
              if (this.currentLevel > this.levels.length) {
                alert("Quiz terminé! Votre score final est: " + this.score);
                this.resetQuiz();
              } else {
                this.showLevelUpMessage();
              }
            }
            this.saveProgress();
          });
    },
    launchConfetti() {
      confetti({
        particleCount: 100,
        spread: 70,
        origin: { y: 0.6 },
        colors: ['#bb0000', '#eb6bff', '00f0fa', '#0000ff', '#f3c13e', '#00ffff', '#ffff00']
      });
    },
    showLevelUpMessage() {
      alert(`Vous passez au niveau ${this.currentLevelName}`);
    },
    resetQuiz() {
      fetch('http://localhost:8081/reset-score', {
        method: 'POST'
      })
          .then(() => {
            this.currentLevel = 1;
            this.currentQuestionIndex = 0;
            this.score = 0;
            this.errorMessage = null;
            this.saveProgress();
          })
          .catch(error => {
            console.error('Error resetting score:', error);
          });
    },
    changePlayer() {
      this.pseudo = '';
      this.resetQuiz();
      localStorage.removeItem('quizPseudo');
    },
    saveProgress() {
      localStorage.setItem('quizCurrentLevel', this.currentLevel);
      localStorage.setItem('quizCurrentQuestionIndex', this.currentQuestionIndex);
      localStorage.setItem('quizScore', this.score);
      localStorage.setItem('quizPseudo', this.pseudo);
    },
    restoreProgress() {
      const savedLevel = localStorage.getItem('quizCurrentLevel');
      const savedQuestionIndex = localStorage.getItem('quizCurrentQuestionIndex');
      const savedScore = localStorage.getItem('quizScore');
      const savedPseudo = localStorage.getItem('quizPseudo');

      if (savedLevel !== null) {
        this.currentLevel = parseInt(savedLevel, 10);
      }
      if (savedQuestionIndex !== null) {
        this.currentQuestionIndex = parseInt(savedQuestionIndex, 10);
      }
      if (savedScore !== null) {
        this.score = parseInt(savedScore, 10);
      }
      if (savedPseudo !== null) {
        this.pseudo = savedPseudo;
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
button {
  margin: 10px;
}

.error-message {
  color: red;
  font-weight: bold;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active in <2.1.8 */ {
  opacity: 0;
}
</style>
