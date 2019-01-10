import Vue from "vue"
import Vuex from "vuex"
import axios from "axios"

Vue.use(Vuex);

const store = () => new Vuex.Store({
  state: {
    name: ""
  },
  mutations: {
    setName (state, name: string) {
      state.name = name;
    }
  }
});

export default store;
