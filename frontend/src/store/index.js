import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    airplaneA: [
      [3, 3, 4, 3, 3, 3, 3, 3],
      [4, 3, 4, 3, 3, 3, 3, 3],
      [4, 4, 4, 5, 3, 3, 3, 3],
      [4, 3, 4, 3, 3, 3, 3, 3],
      [3, 3, 4, 3, 3, 5, 3, 3],
      [3, 3, 3, 4, 4, 4, 4, 4],
      [3, 3, 3, 3, 3, 4, 3, 3],
      [3, 3, 3, 3, 4, 4, 4, 3]
    ],
    airplaneB: [
      [3, 3, 3, 3, 3, 3, 3, 3, 4, 3],
      [3, 3, 3, 3, 3, 3, 3, 4, 3, 3],
      [3, 3, 3, 3, 3, 3, 4, 3, 3, 4],
      [3, 3, 3, 3, 3, 5, 4, 4, 4, 4],
      [3, 3, 3, 3, 3, 3, 4, 3, 3, 4],
      [3, 3, 3, 5, 3, 3, 3, 4, 3, 3],
      [3, 3, 4, 4, 4, 3, 3, 3, 4, 3],
      [3, 4, 3, 4, 3, 4, 3, 3, 3, 3],
      [4, 3, 3, 4, 3, 3, 4, 3, 3, 3],
      [3, 3, 4, 4, 4, 3, 3, 3, 3, 3]
    ],
    rotatePuzzle4Example: [
      ["yellow", "green", "green", "green"],
      ["blue", "blue", "red", "red"],
      ["yellow", "red", "red", "yellow"],
      ["blue", "blue", "yellow", "green"]
    ],
    rotatePuzzle4Target: [
      ["green", "green", "blue", "blue"],
      ["green", "green", "blue", "blue"],
      ["yellow", "yellow", "red", "red"],
      ["yellow", "yellow", "red", "red"]
    ],
    rotatePuzzle6Example: [
      ["red", "green", "yellow", "green", "blue", "yellow"],
      ["yellow", "red", "yellow", "green", "green", "yellow"],
      ["red", "blue", "blue", "red", "blue", "red"],
      ["yellow", "red", "yellow", "red", "blue", "red"],
      ["blue", "yellow", "green", "yellow", "blue", "blue"],
      ["red", "green", "green", "green", "green", "blue"]
    ],
    rotatePuzzle6Target: [
      ["red", "red", "red", "green", "green", "green"],
      ["red", "red", "red", "green", "green", "green"],
      ["red", "red", "red", "green", "green", "green"],
      ["blue", "blue", "blue", "yellow", "yellow", "yellow"],
      ["blue", "blue", "blue", "yellow", "yellow", "yellow"],
      ["blue", "blue", "blue", "yellow", "yellow", "yellow"]
    ],
    rotatePuzzle11Example: [
      [
        "",
        "",
        "purple",
        "light-green",
        "purple",
        "pink",
        "yellow",
        "purple",
        "green",
        "",
        ""
      ],
      [
        "",
        "pink",
        "pink",
        "light-green",
        "green",
        "yellow",
        "yellow",
        "purple",
        "purple",
        "light-green",
        ""
      ],
      [
        "purple",
        "green",
        "pink",
        "yellow",
        "purple",
        "green",
        "purple",
        "orange",
        "purple",
        "orange",
        "light-green"
      ],
      [
        "yellow",
        "yellow",
        "yellow",
        "orange",
        "orange",
        "light-green",
        "pink",
        "light-green",
        "yellow",
        "pink",
        "orange"
      ],
      [
        "",
        "green",
        "green",
        "light-green",
        "orange",
        "green",
        "pink",
        "yellow",
        "pink",
        "green",
        ""
      ],
      [
        "",
        "",
        "orange",
        "green",
        "orange",
        "light-green",
        "pink",
        "orange",
        "light-green",
        "",
        ""
      ]
    ],
    rotatePuzzle11Target: [
      [
        "",
        "",
        "green",
        "light-green",
        "light-green",
        "light-green",
        "light-green",
        "light-green",
        "pink",
        "",
        ""
      ],
      [
        "",
        "green",
        "green",
        "green",
        "light-green",
        "light-green",
        "light-green",
        "pink",
        "pink",
        "pink",
        ""
      ],
      [
        "green",
        "green",
        "green",
        "green",
        "green",
        "light-green",
        "pink",
        "pink",
        "pink",
        "pink",
        "pink"
      ],
      [
        "yellow",
        "yellow",
        "yellow",
        "yellow",
        "yellow",
        "purple",
        "orange",
        "orange",
        "orange",
        "orange",
        "orange"
      ],
      [
        "",
        "yellow",
        "yellow",
        "yellow",
        "purple",
        "purple",
        "purple",
        "orange",
        "orange",
        "orange",
        ""
      ],
      [
        "",
        "",
        "yellow",
        "purple",
        "purple",
        "purple",
        "purple",
        "purple",
        "orange",
        "",
        ""
      ]
    ],
    roomid: ""
  },
  mutations: {
    getRoom: (state, data) => {
      state.roomid = data.roomid;
    }
  }
});
