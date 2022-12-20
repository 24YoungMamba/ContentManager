<script setup></script>

<template>
  <main>
    <h1>Test</h1>

    <img src="./1.svg" />
  </main>
</template>

<style scoped lang="scss">
main {
  width: 80vw;
  height: 100vh;

  /* position: fixed; */
}

h1 {
  font-size: 3em;

  margin: 1vw;
}

img {
  width: 5vw;
  margin: 1vw;
}

// Responsive
@media only screen and (max-width: 1024px) {
} // Laptop + Tablet

@media only screen and (max-width: 600px) {
  main {
    width: 150vw;
  }
} // Mobile
</style>
