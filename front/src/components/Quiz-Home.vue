<template>
  <div>
    <h1>Quiz Aventure</h1>
    <p>Bienvenue dans Quiz Aventure !</p>
    <p>Score: {{ score }}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      score: 0
    };
  }
};
</script>

<style scoped>
h1 {
  color: #2c3e50;
}
</style>
