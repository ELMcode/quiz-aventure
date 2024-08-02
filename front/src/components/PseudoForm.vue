<template>
  <div>
    <h1>Bienvenue dans Quiz Aventure</h1>

    <h2>Règles du jeu :</h2>

    <ul>
      <li>Chaque bonne réponse vous donne 10 points</li>
      <li>Chaque mauvaise réponse vous retire 5 points</li>
      <li>Il y a 5 questions par niveau</li>
      <li>Il y a 3 niveaux : Débutant, Intermédiaire, Expert !</li>
      <li>Vous passez au niveau suivant après avoir répondu à toutes les questions du niveau actuel</li>
    </ul>
    <br>
    <p>Veuillez entrer votre pseudo pour commencer le jeu !</p>

    <form @submit.prevent="submit">
      <input v-model="pseudo" placeholder="Pseudo" required />
      <button type="submit">Jouer</button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      pseudo: ''
    };
  },
  methods: {
    submit() {
      console.log('Emitting pseudo:', this.pseudo);
      this.$emit('submit', this.pseudo);
    }
  }
};
</script>

<style scoped>
h1 {
  color: #2c3e50;
}
form {
  display: flex;
  flex-direction: column;
  align-items: center;
}
input {
  margin-bottom: 10px;
  padding: 10px;
  font-size: 16px;
}
button {
  padding: 10px 20px;
  font-size: 16px;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  margin-bottom: 10px;
  font-size: 16px;
}
</style>
