import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import FindAirplaneIntro from "../views/FindAirplaneIntro.vue";
import FindAirplaneGame from "../views/FindAirplaneGame.vue";
import FindAirplaneLocal from "../views/FindAirplaneLocal.vue";
import TicTacToeBoxIntro from "../views/TicTacToeBoxIntro.vue";
import TicTacToeBoxGame from "../views/TicTacToeBoxGame.vue";
import TicTacToeBoxLocal from "../views/TicTacToeBoxLocal.vue";
import RotatingPuzzleIntro from "../views/RotatingPuzzleIntro.vue";
import RotatingPuzzleLocal from "../views/RotatingPuzzleLocal.vue";
import RotatingPuzzleLocalHard from "../views/RotatingPuzzleLocalHard.vue";
import NotFound from "../views/NotFound.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/Game",
    name: "Home",
    component: Home,
    alias: "/index.html"
  },
  // {
  //   path: "/about",
  //   name: "About",
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () =>
  //     import(/* webpackChunkName: "about" */ "../views/About.vue")
  // },
  {
    path: "/Game/FindAirplane/Introduction",
    name: "FindAirplaneIntroduction",
    component: FindAirplaneIntro
  },
  {
    path: "/Game/FindAirplane/Game/room",
    name: "FindAirplaneGame",
    component: FindAirplaneGame,
    props: route => ({ roomid: route.query.room })
  },
  {
    path: "/Game/FindAirplane/Localgame",
    name: "FindAirplaneLocal",
    component: FindAirplaneLocal,
    props: route => ({ boardlength: route.query.boardlength })
  },
  {
    path: "/Game/TicTacToeBox/Introduction",
    name: "TicTacToeBoxIntroduction",
    component: TicTacToeBoxIntro
  },
  {
    path: "/Game/TicTacToeBox/Game/room",
    name: "TicTacToeBoxGame",
    component: TicTacToeBoxGame,
    props: router => ({ roomid: router.query.room })
  },
  {
    path: "/Game/TicTacToeBox/Localgame",
    name: "TicTacToeBoxLocal",
    component: TicTacToeBoxLocal
  },
  {
    path: "/Game/RotatingPuzzle/Introduction",
    name: "RotatingPuzzleIntroduction",
    component: RotatingPuzzleIntro
  },
  {
    path: "/Game/RotatingPuzzle/Localgame",
    name: "RotatingPuzzleLocal",
    component: RotatingPuzzleLocal,
    props: router => ({ boardtype: router.query.board })
  },
  {
    path: "/Game/RotatingPuzzle/Localgame/Hard",
    name: "RotatingPuzzleLocalHard",
    component: RotatingPuzzleLocalHard
  },
  {
    path: "*",
    name: "NotFound",
    component: NotFound
  }
];

const router = new VueRouter({
  mode: "history",
  routes
});

export default router;
