<template>
  <div class="game-page">
    <h1>
      <font-awesome-icon :icon="['fas', 'plane']" class="icon__airplane title__icon"/>
      <router-link :to="{name: 'Home'}" class="link__none-style">Seek for airplane's head</router-link>
    </h1>

    <notice v-bind:thisPlayer="thisPlayer" v-bind:player2="player2"></notice>

    <div class="game-content">
      Round {{ round }}
      <span class="game-player">&#10142; {{ playerNow }}</span>
      <br>
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

      <!-- Here is the game -->
      <div class="competition">
        <!-- Here is the left part for player 1 -->
        <div>
          <div v-if="thisPlayer == 'player1'" class="game__underline">
            <span class="game-icon__pink">&#10048;</span> Player 1
          </div>
          <div v-else>Player 1</div>
          <!-- airplane board -->
          <table>
            <tr v-for="(array, indexarray) in player1Board" :key="indexarray">
              <td v-for="(block, indexblock) in array" :key="indexblock" class="board-td">
                <div v-if="block === 3" class="airplane__white"></div>
                <div v-else-if="block === 4" class="airplane__blue"></div>
                <div v-else-if="block === 5" class="airplane__darkblue"></div>
                <div v-else-if="thisPlayer == 'player2'" class="airplane__gray"></div>
                <div v-else class="airplane__gray" v-on:click="flip(indexarray, indexblock)"></div>
              </td>
            </tr>
          </table>
        </div>

        <!-- Here is the right part for player 2 -->
        <div>
          <div v-if="thisPlayer == 'player2'" class="game__underline">
            <span class="game-icon__pink">&#10048;</span> Player 2
          </div>
          <div v-else>Player 2</div>
          <!-- airplane board -->
          <table>
            <tr v-for="(array, indexarray) in player2Board" :key="indexarray">
              <td v-for="(block, indexblock) in array" :key="indexblock" class="board-td">
                <div v-if="block === 3" class="airplane__white"></div>
                <div v-else-if="block === 4" class="airplane__blue"></div>
                <div v-else-if="block === 5" class="airplane__darkblue"></div>
                <div v-else-if="thisPlayer == 'player1'" class="airplane__gray"></div>
                <div v-else class="airplane__gray" v-on:click="flip(indexarray, indexblock)"></div>
              </td>
            </tr>
          </table>
        </div>
      </div>
    </div>

    <!-- Here is the example -->
    <div class="intro">
      <airplaneA class="inline-block" v-if="shape == 'airplaneA'"></airplaneA>
      <airplaneB class="inline-block" v-else-if="shape == 'airplaneB'"></airplaneB>

      <div class="inline-block">
        <div class="inline-block">&larr; This is an example of airplanes.</div>
        <br>
        <button class="inline-block introduction-button" v-on:click="clearTime()">
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
import airplaneA from "../components/airplaneA.vue";
import airplaneB from "../components/airplaneB.vue";

export default {
  props: {
    roomid: String,
    shape: String
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
      timer: null
    };
  },
  components: {
    airplaneA,
    airplaneB
  },
  computed: {
    airplaneA() {
      return this.$store.state.airplaneA;
    },
    airplaneB() {
      return this.$store.state.airplaneB;
    }
  },
  mounted: function() {
    this.reload();
  },
  methods: {
    reload() {
      this.axios
        .get(`${this.$hostname}/Game/api/FindAirplane/Game/room`, {
          withCredentials: true,
          params: {
            room: this.roomid
          }
        })
        .then(response => {
          if (response.data.Err != "") {
            window.clearTimeout(this.timer);
            if (response.data.Err === "Sorry! The room is not existing!") {
              this.$router.push({
                path: "/Game/FindAirplane/Introduction"
              });
            }
            if (response.data.Err === "Sorry! This room is full!") {
              console.log("room full");
            }
          } else {
            this.player1Board = response.data.Board1;
            this.player2Board = response.data.Board2;
            this.playerNow = response.data.PlayerNow;
            this.round = response.data.Round;
            this.win = response.data.Win;
            this.thisPlayer = response.data.ThisPlayer;
            this.player1 = response.data.Player1;
            this.player2 = response.data.Player2;
            this.timer = window.setTimeout(this.reload, 1000);
          }
        })
        .catch(err => {
          this.errors.push(err);
        });
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
            `${this.$hostname}/Game/api/FindAirplane/Game/room`,
            {
              Board1: this.player1Board,
              Board2: this.player2Board
            },
            {
              withCredentials: true,
              params: {
                room: this.roomid
              }
            }
          )
          .catch(err => {
            this.errors.push(err);
          });
      }

      this.clearTime();
      this.reload();
    },
    clearTime() {
      window.clearTimeout(this.timer);
    }
  }
};
</script>

<style scoped>
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

@media (max-width: 700px) {
  .competition {
    display: block;
  }
  .board-td {
    width: 35px;
    height: 35px;
  }
}

@media (max-width: 440px) {
  .inline-block {
    display: block;
    margin-right: 30px;
  }
}
</style>
