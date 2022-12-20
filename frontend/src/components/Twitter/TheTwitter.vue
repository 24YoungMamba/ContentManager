<script setup>
import SideNav from "../Design/SideNav.vue";
import Layout from "./Layout/Layout.vue";
</script>

<template>
  <div class="container">
    <SideNav />
    <Layout />
  </div>
</template>

<style scoped lang="scss">
.container {
  display: flex;

  flex-direction: row;
  justify-content: flex-start;
  align-items: flex-start;
  align-content: flex-start;
}

// Responsive
@media only screen and (max-width: 1024px) {
} // Laptop + Tablet

@media only screen and (max-width: 600px) {
  .container {
    flex-direction: row;
    flex-wrap: wrap-reverse;

    justify-content: center;
    align-content: center;
  }
} // Mobile
</style>

<!-- 
  Side Nav

  Main Component
    Welcome Back, @ + name + image
    Last Tweet, with stats
    Weekly Stats
    To-Do
 -->
