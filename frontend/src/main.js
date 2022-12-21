import { createApp } from "vue";
import { createStore } from "vuex";
import createPersistedState from "vuex-persistedstate";

import App from "./App.vue";
import "./style.css";

const store = createStore({
  state: {
    view: "DashboardView",
  },

  mutations: {
    CHANGE_VIEW(state, view) {
      state.view = view;

      console.log(view);
      location.reload();
    },
  },

  plugins: [createPersistedState()],
});
const app = createApp(App);

app.use(store);
app.mount("#app");
