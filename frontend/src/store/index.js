import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    exampleMatrix: [
      [3, 3, 4, 3, 3, 3, 3, 3],
      [4, 3, 4, 3, 3, 3, 3, 3],
      [4, 4, 4, 5, 3, 3, 3, 3],
      [4, 3, 4, 3, 3, 3, 3, 3],
      [3, 3, 4, 3, 3, 5, 3, 3],
      [3, 3, 3, 4, 4, 4, 4, 4],
      [3, 3, 3, 3, 3, 4, 3, 3],
      [3, 3, 3, 3, 4, 4, 4, 3]
    ],
    roomid: ""
  },
  mutations: {
    getRoom: (state, data) => {
      state.roomid = data.roomid;
    }
  }
});
