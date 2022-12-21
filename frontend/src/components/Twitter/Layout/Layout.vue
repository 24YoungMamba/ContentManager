<script setup>
import { ref } from "vue";
import { useStore } from "vuex";

import Dashboard from "../Apps/Dashboard.vue";
import Analytics from "../Apps/Analytics.vue";

const store = useStore();
const view = ref(store.state.view);
</script>

<template>
  <main>
    <Dashboard v-if="view == 'DashboardView'" />
    <Analytics v-else-if="view == 'AnalyticsView'" />
  </main>
</template>

<style scoped lang="scss">
main {
  width: 80vw;
  height: 100vh;
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
    width: 100vw;
  }
} // Mobile
</style>
