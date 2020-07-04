<template>
  <div class="page">
    <h1>
      <font-awesome-icon :icon="['fas', 'plane']" class="icon__airplane"/>Seek for airplane's head
    </h1>

    <div class="notice">
      <font-awesome-icon :icon="['fas', 'bullhorn']" class="icon__bullborn"/>
      <span v-if="thisPlayer == 'player1'">
        Welcome! You are
        <strong>player1</strong> !
      </span>
      <span v-else>Player1 entered the room.</span>
      <span v-if="player2 != ''">
        <span v-if="thisPlayer == 'player2'">
          Welcome! You are
          <strong>player2</strong> !
        </span>
        <span v-else>Player2 entered the room.</span>
      </span>
    </div>

    <div class="game-round-player">
      Round {{ round }}
      <span class="game-player">&#10142; {{ playerNow }}</span>
    </div>
    <div class="game-round-player">
      <div v-for="(winner, indexwinner) in win" :key="indexwinner">
        <div
          v-if="winner == 'player1'"
          class="game-congratulation"
        >&#10047; Congratulate on Player 1!!!</div>
        <div
          v-if="winner == 'player2'"
          class="game-congratulation"
        >&#10047; Congratulate on Player 2!!!</div>
      </div>
    </div>

    <!-- Here is the game -->
    <div class="competition">
      <!-- Here is the left part for player 1 -->
      <div>
        <div v-if="thisPlayer == 'player1'" class="game-round-player game__underline">
          <span class="game-icon__pink">&#10048;</span> Player 1
        </div>
        <div v-else class="game-round-player">Player 1</div>
        <!-- airplane board -->
        <table>
          <tr v-for="(array, indexarray) in player1Board" :key="indexarray">
            <td v-for="(block, indexblock) in array" :key="indexblock" class="board-td">
              <div v-if="block === 3" class="board-td__white"></div>
              <div v-else-if="block === 4" class="board-td__blue"></div>
              <div v-else-if="block === 5" class="board-td__darkblue"></div>
              <div v-else-if="thisPlayer == 'player2'" class="board-td__gray"></div>
              <div v-else class="board-td__gray" v-on:click="flip(indexarray, indexblock)"></div>
            </td>
          </tr>
        </table>
      </div>

      <!-- Here is the right part for player 2 -->
      <div>
        <div v-if="thisPlayer == 'player2'" class="game-round-player game__underline">
          <span class="game-icon__pink">&#10048;</span> Player 2
        </div>
        <div v-else class="game-round-player">Player 2</div>
        <!-- airplane board -->
        <table>
          <tr v-for="(array, indexarray) in player2Board" :key="indexarray">
            <td v-for="(block, indexblock) in array" :key="indexblock" class="board-td">
              <div v-if="block === 3" class="board-td__white"></div>
              <div v-else-if="block === 4" class="board-td__blue"></div>
              <div v-else-if="block === 5" class="board-td__darkblue"></div>
              <div v-else-if="thisPlayer == 'player1'" class="board-td__gray"></div>
              <div v-else class="board-td__gray" v-on:click="flip(indexarray, indexblock)"></div>
            </td>
          </tr>
        </table>
      </div>
    </div>

    <!-- Here is the example -->
    <div class="intro">
      <table class="inline-block">
        <tr v-for="(array, indexarray) in exampleArrays" :key="indexarray">
          <td v-for="(block, indexblock) in array" :key="indexblock" class="board-example">
            <div v-if="block === 3" class="board-td__white"></div>
            <div v-else-if="block === 4" class="board-td__blue"></div>
            <div v-else-if="block === 5" class="board-td__darkblue"></div>
            <div v-else class="board-td__gray"></div>
          </td>
        </tr>
      </table>

      <div class="inline-block">
        <div class="inline-block">&larr; This is an example of airplanes.</div>
        <br>
        <button class="inline-block" v-on:click="clearTime()">
          <router-link :to="{name: 'FindAirplaneIntroduction'}" class="link__none-style">
            Back to introduction page to
            <strong>restart</strong>.
          </router-link>
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    roomid: {
      type: String
    }
  },
  data: function() {
    return {
      round: 0,
      playerNow: "",
      win: [],
      player1Board: [],
      player2Board: [],
      thisPlayer: "",
      player1: "",
      player2: "",
      timeStop: null
    };
  },
  computed: {
    exampleArrays() {
      return this.$store.state.exampleMatrix;
    },
    getRoom() {
      return this.$store.state.roomid;
    }
  },
  mounted: function() {
    this.reload();
  },
  methods: {
    reload() {
      this.$store.commit("getRoom", { roomid: this.roomid });
      this.axios
        .get(`${this.$hostname}/api/FindAirplane/Game/room`, {
          withCredentials: true,
          params: {
            room: this.getRoom
          }
        })
        .then(response => {
          this.player1Board = response.data.Board1;
          this.player2Board = response.data.Board2;
          this.playerNow = response.data.PlayerNow;
          this.round = response.data.Round;
          this.win = response.data.Win;
          this.thisPlayer = response.data.ThisPlayer;
          this.player1 = response.data.Player1;
          this.player2 = response.data.Player2;
        })
        .catch(err => {
          this.errors.push(err);
        });
      this.timeStop = window.setTimeout(this.reload, 1000);
    },
    flip(row, column) {
      if (this.thisPlayer == this.playerNow) {
        if (this.thisPlayer == "player1") {
          this.player1Board[row][column] = this.player1Board[row][column] + 3;
        } else {
          this.player2Board[row][column] = this.player2Board[row][column] + 3;
        }
        this.axios
          .post(
            `${this.$hostname}/api/FindAirplane/Game/room`,
            {
              Board1: this.player1Board,
              Board2: this.player2Board,
              PlayerNow: this.playerNow,
              Round: this.round,
              Win: this.win,
              ThisPlayer: this.thisPlayer,
              Player1: this.player1,
              Player2: this.player2
            },
            {
              withCredentials: true,
              params: {
                room: this.getRoom
              }
            }
          )
          .catch(err => {
            this.errors.push(err);
          });
        this.reload();
      }
    },
    clearTime() {
      clearTimeout(this.timeStop);
    }
  }
};
</script>

<style scoped>
.notice {
  background-color: #c6e2ff;
  background-image: linear-gradient(to right, #c6e2ff, white);
  margin-bottom: 40px;
}

.game-round-player {
  width: 100%;
  text-align: center;
  font-family: "Aladin", cursive;
  font-size: 1.25em;
  font-weight: bold;
  margin-bottom: 10px;
}
.game-player {
  font-size: 1em;
  font-weight: normal;
}
.game-congratulation {
  color: red;
}
.game__underline {
  text-decoration: rgba(139, 106, 247, 0.4) underline;
}
.game-icon__pink {
  color: rgb(139, 106, 247);
}

.competition {
  display: grid;
  display: -ms-grid;
  grid-template-columns: 45% 45%;
  grid-column-gap: 10%;
}
.board-td {
  width: 40px;
  height: 40px;
  background-color: darkgray;
  border: 1px solid black;
  border-collapse: collapse;
}

.intro {
  margin-top: 40px;
  margin-left: auto;
  margin-right: auto;
  width: max-content;
}
.inline-block {
  display: inline-block;
  vertical-align: middle;
  margin-right: 30px;
}
.board-example {
  width: 15px;
  height: 15px;
  background-color: darkgray;
  border: 1px solid black;
  border-collapse: collapse;
}
span {
  padding-right: 5px;
}
button {
  font-family: "Neucha", sans-serif;
  font-size: 1em;
  margin: 15px 0;
  padding: 4px 10px 0 10px;
  border-radius: 3px;
  color: black;
}
button:hover {
  background-color: #003bac;
  border-top: 2px solid #608cdf;
  border-left: 2px solid #608cdf;
  border-bottom: 2px solid #002a7b;
  border-right: 2px solid #002a7b;
}
button:hover a {
  color: white;
}
</style>
