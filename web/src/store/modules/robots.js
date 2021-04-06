import axios from "axios";

export default {
  namespaced: true,
  state: {
    cart: [],
    parts: null,
    foo: "robots-foo"
  },
  mutations: {
    addRobotToCart(state, robot) {
      state.cart.push(robot);
    },
    updateParts(state, parts) {
      state.parts = parts;
    }
  },
  actions: {
    getParts({ commit }) {
      console.log(process.env.VUE_APP_API_URL);
      axios
        .get(`${process.env.VUE_APP_API_URL}/api/parts`)
        .then((result) => commit("updateParts", result.data))
        .catch(console.error);
    },
    addRobotToCart({ commit, state }, robot) {
      const cart = [...state.cart, robot];
      return axios
        .post(`${process.env.VUE_APP_API_URL}/api/cart`, cart)
        .then(() => commit("addRobotToCart", robot))
        .catch(console.error);
    }
  },
  getters: {
    cartSaleItems(state) {
      return state.cart.filter((item) => item.head.onSale);
    },
    foo(state) {
      return `robots-getter/${state.foo}`;
    }
  }
};
