import axios from "axios";

export default {
  state: {
    user: null
  },
  mutations: {
    updateCurrentUser(state, user) {
      state.user = user;
    }
  },
  getters: {
    foo(state, getters, rootState) {
      // getters - non namespaced module, this is all the getters,
      // in a namespaced module these getters are the local getters
      // rootState -
      return `users-getter/${rootState.foo}`;
    }
  },
  actions: {
    signIn({ commit }) {
      axios
        .post("http://localhost:8085/api/sign-in")
        .then((result) => commit("updateCurrentUser", result.data))
        .catch(console.error);
    }
  }
};
