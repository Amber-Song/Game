<template>
  <div class="game-page">
    <h1>
      <font-awesome-icon :icon="['fas', 'plane']" class="icon__airplane title__icon"/>
      <router-link :to="{name: 'Home'}" class="link__none-style">Seek for airplane's head</router-link>
    </h1>

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
          <div>Player 1</div>
          <!-- airplane board -->
          <table>
            <tr v-for="(array, indexarray) in player1Board" :key="indexarray">
              <td v-for="(block, indexblock) in array" :key="indexblock" class="board-td">
                <div v-if="block === 3" class="airplane__white"></div>
                <div v-else-if="block === 4" class="airplane__blue"></div>
                <div v-else-if="block === 5" class="airplane__darkblue"></div>
                <div
                  v-else
                  class="airplane__gray"
                  v-on:click="flip('player1',indexarray, indexblock)"
                ></div>
              </td>
            </tr>
          </table>
        </div>

        <!-- Here is the right part for player 2 -->
        <div>
          <div>Player 2</div>
          <!-- airplane board -->
          <table>
            <tr v-for="(array, indexarray) in player2Board" :key="indexarray">
              <td v-for="(block, indexblock) in array" :key="indexblock" class="board-td">
                <div v-if="block === 3" class="airplane__white"></div>
                <div v-else-if="block === 4" class="airplane__blue"></div>
                <div v-else-if="block === 5" class="airplane__darkblue"></div>
                <div
                  v-else
                  class="airplane__gray"
                  v-on:click="flip('player2',indexarray, indexblock)"
                ></div>
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
        <button class="inline-block introduction-button">
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
    boardlength: Number,
    shape: String
  },
  data: function() {
    return {
      player1Board: [],
      player2Board: [],
      playerNow: "player1",
      round: 1,
      win: []
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
    this.player1Board = this.generateAirplane();
    this.player2Board = this.generateAirplane();
  },
  methods: {
    flip(player, row, column) {
      if (player == "player1" && this.playerNow == "player1") {
        this.player1Board[row][column] = this.player1Board[row][column] + 3;
        this.playerNow = "player2";
        if (!this.win.includes("player1") && this.checkWin(this.player1Board)) {
          this.win[this.win.length] = "player1";
        }
      }
      if (player == "player2" && this.playerNow == "player2") {
        this.player2Board[row][column] = this.player2Board[row][column] + 3;
        this.playerNow = "player1";
        this.round = this.round + 1;
        if (!this.win.includes("player2") && this.checkWin(this.player2Board)) {
          this.win[this.win.length] = "player2";
        }
      }
    },

    checkWin(board) {
      var airplaneFoundNum = 0;
      for (let i = 0; i < this.boardlength; i++) {
        for (let j = 0; j < this.boardlength; j++) {
          if (board[i][j] == 5) {
            airplaneFoundNum = airplaneFoundNum + 1;
          }
        }
      }
      if (airplaneFoundNum == 2) {
        return true;
      }
      return false;
    },

    generateAirplane() {
      // initiate boards
      var board = this.iniBoard();

      var i = 0;
      var j = 0;
      var direction = "";
      var airplane2Find = false;

      if (this.shape == "airplaneA") {
        do {
          // redo playerboard
          board = this.iniBoard();
          // Find airplane1
          do {
            [i, j] = this.generateHead();
            direction = this.generateDirection();
          } while (!this.checkAvailableShapeA(i, j, direction));
          // place the first airplane
          board = this.placeAirplaneA(board, i, j, direction);

          // Try to find airplane2
          airplane2Find = false;
          // Try 5 times to find airplane2,
          //  If still can't, redo airplane1
          for (let checkTime = 0; checkTime < 5; checkTime++) {
            [i, j] = this.generateHead();
            // Try different direction
            for (let index = 0; index < 4; index++) {
              direction = this.generateDirection();
              if (this.checkAvailableShapeA(i, j, direction)) {
                if (this.checkOverlapShapeA(board, i, j, direction)) {
                  airplane2Find = true;
                  break;
                }
              }
            }

            // If airplane2 found, break try to find airplane2 loop
            if (airplane2Find) {
              break;
            }
          }

          // If airplane2 found, then break the regenerate
          // else redo the airplane
        } while (!airplane2Find);

        // Place head 2
        board = this.placeAirplaneA(board, i, j, direction);

        return board;
      }

      if (this.shape == "airplaneB") {
        do {
          // redo playerboard
          board = this.iniBoard();
          // Find airplane1
          do {
            [i, j] = this.generateHead();
            direction = this.generateDirection();
          } while (!this.checkAvailableShapeB(i, j, direction));
          // place the first airplane
          board = this.placeAirplaneB(board, i, j, direction);

          // Try to find airplane2
          airplane2Find = false;
          // Try 5 times to find airplane2,
          //  If still can't, redo airplane1
          for (let checkTime = 0; checkTime < 5; checkTime++) {
            [i, j] = this.generateHead();
            // Try different direction
            for (let index = 0; index < 4; index++) {
              direction = this.generateDirection();
              if (this.checkAvailableShapeB(i, j, direction)) {
                if (this.checkOverlapShapeB(board, i, j, direction)) {
                  airplane2Find = true;
                  break;
                }
              }
            }

            // If airplane2 found, break try to find airplane2 loop
            if (airplane2Find) {
              break;
            }
          }

          // If airplane2 found, then break the regenerate
          // else redo the airplane
        } while (!airplane2Find);

        // Place head 2
        board = this.placeAirplaneB(board, i, j, direction);

        return board;
      }
    },

    placeAirplaneA(board, i, j, direction) {
      board[i][j] = 2;

      switch (direction) {
        case "up":
          for (let index = i + 1; index <= i + 3; index++) {
            board[index][j] = 1;
          }
          for (let index = j - 2; index <= j + 2; index++) {
            board[i + 1][index] = 1;
          }
          for (let index = j - 1; index <= j + 1; index++) {
            board[i + 3][index] = 1;
          }
          return board;
        case "down":
          for (let index = i - 1; index >= i - 3; index--) {
            board[index][j] = 1;
          }
          for (let index = j - 2; index <= j + 2; index++) {
            board[i - 1][index] = 1;
          }
          for (let index = j - 1; index <= j + 1; index++) {
            board[i - 3][index] = 1;
          }
          return board;
        case "left":
          for (let index = j + 1; index <= j + 3; index++) {
            board[i][index] = 1;
          }
          for (let index = i - 2; index <= i + 2; index++) {
            board[index][j + 1] = 1;
          }
          for (let index = i - 1; index <= i + 1; index++) {
            board[index][j + 3] = 1;
          }
          return board;
        case "right":
          for (let index = j - 1; index >= j - 3; index--) {
            board[i][index] = 1;
          }
          for (let index = i - 2; index <= i + 2; index++) {
            board[index][j - 1] = 1;
          }
          for (let index = i - 1; index <= i + 1; index++) {
            board[index][j - 3] = 1;
          }
          return board;
        default:
          break;
      }
      return board;
    },

    checkOverlapShapeA(board, i, j, direction) {
      if (board[i][j] != 0) {
        return false;
      }
      switch (direction) {
        case "up":
          for (let index = i + 1; index <= i + 3; index++) {
            if (board[index][j] != 0) {
              return false;
            }
          }
          for (let index = j - 2; index <= j + 2; index++) {
            if (board[i + 1][index] != 0) {
              return false;
            }
          }
          for (let index = j - 1; index <= j + 1; index++) {
            if (board[i + 3][index] != 0) {
              return false;
            }
          }
          return true;
        case "down":
          for (let index = i - 1; index >= i - 3; index--) {
            if (board[index][j] != 0) {
              return false;
            }
          }
          for (let index = j - 2; index <= j + 2; index++) {
            if (board[i - 1][index] != 0) {
              return false;
            }
          }
          for (let index = j - 1; index <= j + 1; index++) {
            if (board[i - 3][index] != 0) {
              return false;
            }
          }
          return true;
        case "left":
          for (let index = j + 1; index <= j + 3; index++) {
            if (board[i][index] != 0) {
              return false;
            }
          }
          for (let index = i - 2; index <= i + 2; index++) {
            if (board[index][j + 1] != 0) {
              return false;
            }
          }
          for (let index = i - 1; index <= i + 1; index++) {
            if (board[index][j + 3] != 0) {
              return false;
            }
          }
          return true;
        case "right":
          for (let index = j - 1; index >= j - 3; index--) {
            if (board[i][index] != 0) {
              return false;
            }
          }
          for (let index = i - 2; index <= i + 2; index++) {
            if (board[index][j - 1] != 0) {
              return false;
            }
          }
          for (let index = i - 1; index <= i + 1; index++) {
            if (board[index][j - 3] != 0) {
              return false;
            }
          }
          return true;
        default:
          break;
      }
      return false;
    },

    checkAvailableShapeA(i, j, direction) {
      switch (direction) {
        case "up":
          if (this.boardlength - i <= 3 || this.boardlength - j <= 2 || j < 2) {
            return false;
          }
          return true;
        case "down":
          if (i < 3 || this.boardlength - j <= 2 || j < 2) {
            return false;
          }
          return true;
        case "left":
          if (this.boardlength - j <= 3 || this.boardlength - i <= 2 || i < 2) {
            return false;
          }
          return true;
        case "right":
          if (j < 3 || this.boardlength - i <= 2 || i < 2) {
            return false;
          }
          return true;
        default:
          return false;
      }
    },

    placeAirplaneB(board, i, j, direction) {
      board[i][j] = 2;

      switch (direction) {
        case "up":
          for (let index = i + 1; index <= i + 4; index++) {
            board[index][j] = 1;
          }
          for (let index = 1; index <= 3; index++) {
            board[i + index][j + index] = 1;
            board[i + index][j - index] = 1;
          }
          for (let index = j - 1; index <= j + 1; index++) {
            board[i + 4][index] = 1;
          }
          return board;
        case "down":
          for (let index = i - 1; index >= i - 4; index--) {
            board[index][j] = 1;
          }
          for (let index = 1; index <= 3; index++) {
            board[i - index][j + index] = 1;
            board[i - index][j - index] = 1;
          }
          for (let index = j - 1; index <= j + 1; index++) {
            board[i - 4][index] = 1;
          }
          return board;
        case "left":
          for (let index = j + 1; index <= j + 4; index++) {
            board[i][index] = 1;
          }
          for (let index = 1; index <= 3; index++) {
            board[i + index][j + index] = 1;
            board[i - index][j + index] = 1;
          }
          for (let index = i - 1; index <= i + 1; index++) {
            board[index][j + 4] = 1;
          }
          return board;
        case "right":
          for (let index = j - 1; index >= j - 4; index--) {
            board[i][index] = 1;
          }
          for (let index = 1; index <= 3; index++) {
            board[i + index][j - index] = 1;
            board[i - index][j - index] = 1;
          }
          for (let index = i - 1; index <= i + 1; index++) {
            board[index][j - 4] = 1;
          }
          return board;
        default:
          break;
      }
      return board;
    },

    checkOverlapShapeB(board, i, j, direction) {
      if (board[i][j] != 0) {
        return false;
      }
      switch (direction) {
        case "up":
          for (let index = i + 1; index <= i + 4; index++) {
            if (board[index][j] != 0) {
              return false;
            }
          }
          for (let index = 1; index <= 3; index++) {
            if (
              board[i + index][j + index] != 0 ||
              board[i + index][j - index] != 0
            ) {
              return false;
            }
          }
          for (let index = j - 1; index <= j + 1; index++) {
            if (board[i + 4][index] != 0) {
              return false;
            }
          }
          return true;
        case "down":
          for (let index = i - 1; index >= i - 4; index--) {
            if (board[index][j] != 0) {
              return false;
            }
          }
          for (let index = 1; index <= 3; index++) {
            if (
              board[i - index][j + index] != 0 ||
              board[i - index][j - index] != 0
            ) {
              return false;
            }
          }
          for (let index = j - 1; index <= j + 1; index++) {
            if (board[i - 4][index] != 0) {
              return false;
            }
          }
          return true;
        case "left":
          for (let index = j + 1; index <= j + 4; index++) {
            if (board[i][index] != 0) {
              return false;
            }
          }
          for (let index = 1; index <= 3; index++) {
            if (
              board[i + index][j + index] != 0 ||
              board[i - index][j + index] != 0
            ) {
              return false;
            }
          }
          for (let index = i - 1; index <= i + 1; index++) {
            if (board[index][j + 4] != 0) {
              return false;
            }
          }
          return true;
        case "right":
          for (let index = j - 1; index >= j - 4; index--) {
            if (board[i][index] != 0) {
              return false;
            }
          }
          for (let index = 1; index <= 3; index++) {
            if (
              board[i + index][j - index] != 0 ||
              board[i - index][j - index] != 0
            ) {
              return false;
            }
          }
          for (let index = i - 1; index <= i + 1; index++) {
            if (board[index][j - 4] != 0) {
              return false;
            }
          }
          return true;
        default:
          break;
      }
      return false;
    },

    checkAvailableShapeB(i, j, direction) {
      switch (direction) {
        case "up":
          if (this.boardlength - i <= 4 || this.boardlength - j <= 3 || j < 3) {
            return false;
          }
          return true;
        case "down":
          if (i < 4 || this.boardlength - j <= 3 || j < 3) {
            return false;
          }
          return true;
        case "left":
          if (this.boardlength - j <= 4 || this.boardlength - i <= 3 || i < 3) {
            return false;
          }
          return true;
        case "right":
          if (j < 4 || this.boardlength - i <= 3 || i < 3) {
            return false;
          }
          return true;
        default:
          return false;
      }
    },

    generateDirection() {
      var direction = this.generateInt(4);
      switch (direction) {
        case 0:
          return "up";
        case 1:
          return "down";
        case 2:
          return "left";
        case 3:
          return "right";
        default:
          break;
      }
    },

    generateHead() {
      var i = this.generateInt(this.boardlength);
      var j = this.generateInt(this.boardlength);
      return [i, j];
    },

    generateInt(num) {
      return Math.floor(Math.random() * num);
    },

    iniBoard() {
      var board = [];
      var line = [];

      for (let i = 0; i < this.boardlength; i++) {
        line[i] = 0;
      }
      for (let i = 0; i < this.boardlength; i++) {
        board[i] = line.slice();
      }

      return board;
    }
  }
};
</script>

<style scoped>
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
