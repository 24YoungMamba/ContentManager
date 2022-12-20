<script setup>
import { ref, computed } from "vue";

import Home from "./components/Home/Home.vue";

import TheTwitter from "./components/Twitter/TheTwitter.vue";

import TheNotFound from "./components/TheNotFound/TheNotFound.vue";

const routes = {
  "/": Home,
  "#/": Home,
  "#/twitter": TheTwitter,
};

const currentPath = ref(window.location.hash);

window.addEventListener("hashchange", () => {
  currentPath.value = window.location.hash;
});

const currentView = computed(() => {
  return routes[currentPath.value || "/"] || TheNotFound;
});
</script>

<template>
  <KeepAlive>
    <component :is="currentView" />
  </KeepAlive>
</template>

<style scoped></style>
