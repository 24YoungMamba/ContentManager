<script setup>
import { reactive } from "vue";

const assets = reactive([
  {
    src: "src/components/Home/Cards/assets/linkedin.svg",
    firstLetter: "L",
    name: "inkedIn",
    url: "#/linkedin",
    color: "#0288D1",
  },
  {
    src: "src/components/Home/Cards/assets/medium.svg",
    firstLetter: "M",
    name: "edium",
    url: "#/medium",
    color: "#000000",
  },
  {
    src: "src/components/Home/Cards/assets/twitter.svg",
    firstLetter: "T",
    name: "witter",
    url: "#/twitter",
    color: "#03A9F4",
  },
  {
    src: "src/components/Home/Cards/assets/youtube.svg",
    firstLetter: "Y",
    name: "ouTube",
    url: "#/youtube",
    color: "#FF3D00",
  },
  {
    src: "src/components/Home/Cards/assets/instagram.svg",
    firstLetter: "I",
    name: "nstagram",
    url: "#/instagram",
    color: "#5D5FC1",
  },
]);

function redirect(url) {
  window.location = url;
}
</script>

<template>
  <div class="card-container">
    <section v-for="asset in assets" class="card" @click="redirect(asset.url)">
      <img :src="asset.src" />

      <div class="card-text">
        <h1 class="title" :style="asset">
          {{ asset.firstLetter }}
        </h1>
        <h2 class="subtitle">{{ asset.name }}</h2>
      </div>
    </section>
  </div>
</template>

<style scoped>
.card-container {
  height: 50%;

  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  align-content: center;
  flex-wrap: wrap;
}

.card {
  display: flex;
  flex-direction: column;
  align-items: center;

  border: 0.5em solid #fff;
  border-radius: 1.5em;

  width: 10vw; /* Mobile: 20vw */
  padding: calc(3em + var(--coef-height)) calc(1em + var(--coef-width));
  margin: 0.5em 1em;

  transition: all 0.4s ease-in-out;
}

.card:hover {
  transform: translateY(-1em) translateX(-1em);
  box-shadow: 0.5em 0.5em 5px 0px rgba(200, 200, 200, 0.6);
}

.card img {
  width: calc(14em + var(--coef-width));
  background-color: rgba(0, 0, 0, 0);
}

.card:hover img {
  animation: image-animation 1s forwards;
}

@keyframes image-animation {
  0% {
    transform: translateY(0) rotate(0) scale(1);
    filter: drop-shadow(0em 0em 0em);
  }
  50% {
    transform: translateY(-30%) rotate(-10deg) scale(1.15);
    filter: drop-shadow(0.2em 0.5em 1em white);
  }
  100% {
    transform: translateY(-30%) rotate(0) scale(1);
    filter: drop-shadow(0.5em 0.1em 1em white);
  }
}

.card .card-text {
  display: flex;
  align-items: baseline;
  justify-content: center;

  margin-top: 1em;
}

.card .card-text .title {
  color: var(color);
  font-size: calc(var(--fs-title) - 0.5em);
}

.card-text .subtitle {
  color: white;
  font-size: calc(var(--fs-subtitle) - 0.5em);
}

@media only screen and (max-width: 1024px) {
  .card-container {
    height: 60%;
  }

  .card-container .card {
    padding: calc(2em + var(--coef-height)) calc(2em + var(--coef-width));
    /* padding: 1.5em 2em; */
  }

  .card-container .card img {
    width: calc(8em + var(--coef-width));
  }

  .card .card-text .title {
    font-size: calc(var(--fs-title) - 1em);
  }

  .card .card-text .subtitle {
    font-size: calc(var(--fs-subtitle) - 1em);
  }
}

@media only screen and (max-width: 440px) {
  .card-container .card {
    padding: calc(1em + var(--coef-height)) calc(3em + var(--coef-width));
  }

  .card img {
    width: calc(7.5em + var(--coef-width));
  }

  .card .card-text .title {
    font-size: calc(var(--fs-title) - 1.5em);
  }

  .card .card-text .subtitle {
    font-size: calc(var(--fs-subtitle) - 1.2em);
  }
}
</style>
