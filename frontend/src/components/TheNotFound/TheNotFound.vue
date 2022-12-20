<script setup></script>

<template>
  <div class="container">
    <div class="container-404">
      <h1 class="text-404">4</h1>
      <img src="./assets/image_404.svg" alt="404" class="icon-404" />
      <h1 class="text-404">4</h1>
    </div>

    <h2>The Page Does Not Exist!</h2>
    <a href="#/"> Go Back ¯\_(ツ)_/¯ </a>
  </div>
</template>

<style scoped>
.container {
  width: 100vw;
  height: 100vh;

  display: flex;
  flex-direction: column;

  align-items: center;
  justify-content: center;
}

.container-404 {
  display: flex;
  flex-direction: row;

  align-items: center;

  margin: calc(10em + 1vh);
}

.icon-404 {
  width: calc(35em + 1vw);
}

.text-404 {
  font-size: calc(20em + 1vw);
  color: var(--color-primary);
}

h2 {
  font-size: calc(0.1em + 1vw);
}

a {
  font-size: calc(0.05em + 1vw);
}

@media only screen and (max-width: 1024px) {
  h2 {
    font-size: calc(0.1em + 2vw);
  }

  a {
    font-size: calc(0.05em + 2vw);
  }
}

@media only screen and (max-width: 768px) {
  .icon-404 {
    width: calc(20em + 2vw);
  }

  .text-404 {
    font-size: calc(10em + 2vw);
  }

  h2 {
    font-size: calc(1em + 1vw);
  }

  a {
    font-size: calc(0.5em + 1vw);
  }
}
</style>
