import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Notice from "../views/Notice.vue";
import FindAirplaneIntro from "../views/FindAirplaneIntro.vue";
//import FindAirplaneGame from "../views/FindAirplaneGame.vue";
//import FindAirplaneLocal from "../views/FindAirplaneLocal.vue";
import TicTacToeBoxIntro from "../views/TicTacToeBoxIntro.vue";
//import TicTacToeBoxGame from "../views/TicTacToeBoxGame.vue";
//import TicTacToeBoxLocal from "../views/TicTacToeBoxLocal.vue";
import RotatingPuzzleIntro from "../views/RotatingPuzzleIntro.vue";
import RotatingPuzzleLocal from "../views/RotatingPuzzleLocal.vue";
import RotatingPuzzleLocalHard from "../views/RotatingPuzzleLocalHard.vue";
import MonopolyIntro from "../views/MonopolyIntro.vue";
//import MonopolyLocal from "../views/MonopolyLocal.vue";
import NotFound from "../views/NotFound.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "Home",
      component: Home,
      alias: "/index.html"
    },
    {
      path: "/notice",
      name: "Notice",
      component: Notice
    },
    {
      path: "/FindAirplane/Introduction",
      name: "FindAirplaneIntroduction",
      component: FindAirplaneIntro
    },
//    {
//      path: "/FindAirplane/Game/room",
//      name: "FindAirplaneGame",
//      component: FindAirplaneGame,
//      props: route => ({ roomid: route.query.room, shape: route.query.shape })
//    },
//    {
//      path: "/FindAirplane/Localgame",
//      name: "FindAirplaneLocal",
//      component: FindAirplaneLocal,
//      props: route => ({
//        boardlength: route.query.boardlength,
//        shape: route.query.shape
//      })
//    },
    {
      path: "/TicTacToeBox/Introduction",
      name: "TicTacToeBoxIntroduction",
      component: TicTacToeBoxIntro
    },
//    {
//      path: "/TicTacToeBox/Game/room",
//      name: "TicTacToeBoxGame",
//      component: TicTacToeBoxGame,
//      props: router => ({ roomid: router.query.room })
//    },
//    {
//      path: "/TicTacToeBox/Localgame",
//      name: "TicTacToeBoxLocal",
//      component: TicTacToeBoxLocal
//    },
    {
      path: "/RotatingPuzzle/Introduction",
      name: "RotatingPuzzleIntroduction",
      component: RotatingPuzzleIntro
    },
    {
      path: "/RotatingPuzzle/Localgame",
      name: "RotatingPuzzleLocal",
      component: RotatingPuzzleLocal,
      props: router => ({ boardtype: router.query.board })
    },
    {
      path: "/RotatingPuzzle/Localgame/Hard",
      name: "RotatingPuzzleLocalHard",
      component: RotatingPuzzleLocalHard
    },
    {
      path: "/Monopoly/Introduction",
      name: "MonopolyIntroduction",
      component: MonopolyIntro
    },
//    {
//      path: "/Monopoly/Localgame",
//      name: "MonopolyLocal",
//      component: MonopolyLocal,
//      props: router => ({ playerNum: router.query.playerNum })
//    },
    {
      path: "/.*",
      name: "NotFound",
      component: NotFound
    }
  ]
})

export default router
