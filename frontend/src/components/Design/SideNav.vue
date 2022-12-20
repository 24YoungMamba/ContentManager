<script setup>
import { ref } from "vue";

const isOpen = ref(false);
</script>

<template>
  <nav :class="isOpen ? 'opened' : ' '">
    <section class="menu">
      <ul>
        <a @click="isOpen = !isOpen">
          <TransitionGroup name="burger-icon">
            <img class="one" src="./assets/burger.svg" v-if="!isOpen" />
            <img class="two" src="./assets/close.svg" v-if="isOpen" />
          </TransitionGroup>
        </a>
      </ul>
    </section>

    <section class="main">
      <p class="helper-text">Main:</p>
      <ul style="height: 50vh">
        <a href="">
          <p v-if="isOpen">Dashboard</p>
          <img src="./assets/dashboard.svg" />
        </a>
        <a href="">
          <p v-if="isOpen">Analytics</p>
          <img src="./assets/chart.svg" />
        </a>
        <a href="">
          <p v-if="isOpen">Weekly Report</p>
          <img src="./assets/week.svg" />
        </a>
        <a href="">
          <p v-if="isOpen">Ideas</p>
          <img src="./assets/ideas.svg" />
        </a>
      </ul>
    </section>

    <section class="others">
      <p class="helper-text">Others:</p>
      <ul>
        <a href="#/">
          <p v-if="isOpen">Go Back</p>
          <img src="./assets/back.svg" style="transform: rotate(180deg)" />
        </a>
      </ul>
    </section>
  </nav>
</template>

<style scoped lang="scss">
// Layout
nav {
  width: 5vw;
  height: 100vh;

  display: flex;

  flex-direction: column;
  justify-content: space-around;
  align-items: center;
  align-content: center;

  background-color: #1f1f1f;

  transition: all 0.5s;

  & ul {
    width: 100%;

    display: flex;

    flex-direction: column;
    align-items: flex-start;

    & * {
      margin-bottom: calc(2vh + var(--coef-height));
    }
  }

  & .menu {
    width: 48px;
    height: 48px;

    z-index: 999;
  }

  & * {
    background-color: #1f1f1f;
  }
}

// Text Styling
a {
  text-decoration: none;

  display: flex;
  flex-direction: row;
  align-items: center;

  font-size: calc(0.6vw + var(--coef-width));

  & p {
    font-size: calc(0.7vw + var(--coef-width));
    margin-right: 1vw;
  }
}

p {
  text-align: center;

  font-size: calc(0.3vw + var(--coef-width));
  font-weight: 100;
}

// Burger Menu
.one {
  position: absolute;
  z-index: 2;

  width: 48px;
}

.two {
  position: absolute;

  width: 48px;
}

.burger-icon-enter-to,
.burger-icon-leave-from {
  animation: fadeIn 0.5s forwards;
}

.burger-icon-enter-from,
.burger-icon-leave-to {
  animation: fadeIn 0.5s reverse;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

// Transition
.opened {
  width: 15vw;

  & ul {
    align-items: flex-end;
  }

  & a {
    transition-delay: 1s;
  }
}

// Responsive
@media only screen and (max-width: 1024px) {
  nav {
    width: 10vw;

    & p {
      font-size: calc(0.8vw + var(--coef-width));
    }
  }

  .opened {
    width: 25vw;

    & ul {
      align-items: flex-end;
    }

    & a {
      & p {
        font-size: calc(1.4vw + var(--coef-width));
      }
    }
  }
} // Laptop + Tablet

@media only screen and (max-width: 600px) {
  nav {
    width: 100vw;
    height: calc(7.5vh + 2%);

    position: fixed;
    display: block;

    & ul {
      flex-direction: row;
      align-items: flex-start;
      justify-content: space-evenly;

      width: 100% !important;
      height: 48px !important;

      margin-top: calc(1vh + 1%);

      & * {
        margin-bottom: 0;
      }
    }

    & .main {
      width: 100%;
    }
    & .menu {
      display: none;
      position: none;
    }
    & .others {
      display: none;
      position: none;
    }

    & p {
      display: none;
    }
  }
} // Mobile
</style>
