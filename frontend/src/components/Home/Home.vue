<script setup>
import TheCards from "./Cards/TheCards.vue";
</script>

<template>
  <div class="container">
    <h1 class="title">Young Mamba's</h1>
    <h4 class="subtitle">Content Manager</h4>

    <TheCards />
  </div>
</template>

<style scoped>
.container {
  width: 100vw;
  height: 100vh;

  display: flex;
  flex-direction: column;

  justify-content: center;

  align-items: center;
  align-content: center;
}

.title {
  font-size: var(--fs-title);
  font-weight: var(--fw-title);
  text-align: center;
}

.subtitle {
  font-size: var(--fs-subtitle);
  font-weight: var(--fw-subtitle);
  text-align: center;
}

@media only screen and (max-width: 440px) {
  .title {
    font-size: calc(var(--fs-title) - 1em);
  }

  .subtitle {
    font-size: calc(var(--fs-subtitle) - 1em);
  }
}
</style>
