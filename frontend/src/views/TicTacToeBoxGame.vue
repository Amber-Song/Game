<template>
  <div class="page">
    <h1>
      <span class="fa fa-cube"></span>
      <router-link :to="{name: 'Home'}" class="link__none-style">Tic Tac Toe Box version</router-link>
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
          <strong>Game start !</strong>
        </span>
        <span v-else>
          Player2 entered the room.
          <strong>Game start !</strong>
        </span>
      </span>
      <span v-else>Please wait for player2 !</span>
    </div>

    <div class="competition">
      <!-- This is player1 collection -->
      <div class="game-round-player">
        player1
        <br>
        <div v-for="(box, index) in collection1" :key="index">
          <div v-if="thisPlayer == 'player1'">
            <div
              v-if="box == 'small'"
              class="fa fa-cube cube__small blue"
              v-on:click="choosebox(index)"
              v-bind:class="{chosen: index === boxChosenIndex}"
            ></div>
            <div
              v-if="box == 'medium'"
              class="fa fa-cube cube__medium blue"
              v-on:click="choosebox(index)"
              v-bind:class="{chosen: index === boxChosenIndex}"
            ></div>
            <div
              v-if="box == 'large'"
              class="fa fa-cube cube__large blue"
              v-on:click="choosebox(index)"
              v-bind:class="{chosen: index === boxChosenIndex}"
            ></div>
          </div>
          <div v-else>
            <div v-if="box == 'small'" class="fa fa-cube cube__small blue"></div>
            <div v-if="box == 'medium'" class="fa fa-cube cube__medium blue"></div>
            <div v-if="box == 'large'" class="fa fa-cube cube__large blue"></div>
          </div>
        </div>
      </div>

      <!-- This is board -->
      <div class="competition__board game-round-player">
        <div>
          Round: {{round}}
          <span class="game-player">&#10142; {{playerNow}}</span>
        </div>
        <div v-if="winner != ''" class="game-congratulation">Congratulations {{winner}} !</div>
        <div>{{notice}}</div>
        <table>
          <tr v-for="(line, indexline) in boardDisplay" :key="indexline">
            <td
              v-for="(cell, indexcell) in line"
              :key="indexcell"
              v-on:click="chooseCell(indexline,indexcell)"
            >
              <div v-if="cell == 'small'">
                <div
                  v-if="boardDisplayPlayer[indexline][indexcell] == 'player1'"
                  class="fa fa-cube cube__small blue"
                  v-on:click="chooseBoxFromBoard(indexline,indexcell)"
                  v-bind:class="{chosen: (indexline === boxChosenI && indexcell === boxChosenJ)}"
                ></div>
                <div
                  v-else
                  class="fa fa-cube cube__small green"
                  v-on:click="chooseBoxFromBoard(indexline,indexcell)"
                  v-bind:class="{chosen: (indexline === boxChosenI && indexcell === boxChosenJ)}"
                ></div>
              </div>
              <div v-else-if="cell == 'medium'">
                <div
                  v-if="boardDisplayPlayer[indexline][indexcell] == 'player1'"
                  class="fa fa-cube cube__medium blue"
                  v-on:click="chooseBoxFromBoard(indexline,indexcell)"
                  v-bind:class="{chosen: (indexline === boxChosenI && indexcell === boxChosenJ)}"
                ></div>
                <div
                  v-else
                  class="fa fa-cube cube__medium green"
                  v-on:click="chooseBoxFromBoard(indexline,indexcell)"
                  v-bind:class="{chosen: (indexline === boxChosenI && indexcell === boxChosenJ)}"
                ></div>
              </div>
              <div v-else-if="cell == 'large'">
                <div
                  v-if="boardDisplayPlayer[indexline][indexcell] == 'player1'"
                  class="fa fa-cube cube__large blue"
                  v-on:click="chooseBoxFromBoard(indexline,indexcell)"
                  v-bind:class="{chosen: (indexline === boxChosenI && indexcell === boxChosenJ)}"
                ></div>
                <div
                  v-else
                  class="fa fa-cube cube__large green"
                  v-on:click="chooseBoxFromBoard(indexline,indexcell)"
                  v-bind:class="{chosen: (indexline === boxChosenI && indexcell === boxChosenJ)}"
                ></div>
              </div>
              <div v-else></div>
            </td>
          </tr>
        </table>
      </div>

      <!-- This is player2 collection -->
      <div class="game-round-player">
        player2
        <br>
        <div v-for="(box, index) in collection2" :key="index">
          <div v-if="thisPlayer == 'player2'">
            <div
              v-if="box == 'small'"
              class="fa fa-cube cube__small green"
              v-on:click="choosebox(index)"
              v-bind:class="{chosen: index === boxChosenIndex}"
            ></div>
            <div
              v-if="box == 'medium'"
              class="fa fa-cube cube__medium green"
              v-on:click="choosebox(index)"
              v-bind:class="{chosen: index === boxChosenIndex}"
            ></div>
            <div
              v-if="box == 'large'"
              class="fa fa-cube cube__large green"
              v-on:click="choosebox(index)"
              v-bind:class="{chosen: index === boxChosenIndex}"
            ></div>
          </div>
          <div v-else>
            <div v-if="box == 'small'" class="fa fa-cube cube__small green"></div>
            <div v-if="box == 'medium'" class="fa fa-cube cube__medium green"></div>
            <div v-if="box == 'large'" class="fa fa-cube cube__large green"></div>
          </div>
        </div>
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
      thisPlayer: "",
      player1: "",
      player2: "",
      collection1: [],
      collection2: [],
      board: [],
      boardPlayer: [],
      winner: "",
      round: 0,
      playerNow: "",
      boardDisplay: [],
      boardDisplayPlayer: [],
      boxChosenIndex: -1,
      boxChosen: "",
      boxChosenI: -1,
      boxChosenJ: -1,
      notice: ""
    };
  },

  mounted: function() {
    // initate boardDisplay
    var board = [];
    var boardPlayer = [];
    var line = [];
    for (let i = 0; i < 3; i++) {
      line[i] = "";
    }
    for (let i = 0; i < 3; i++) {
      board[i] = line.slice();
      boardPlayer[i] = line.slice();
    }
    this.boardDisplay = board.slice();
    this.boardDisplayPlayer = boardPlayer.slice();

    this.$store.commit("getRoom", { roomid: this.roomid });
    this.loadData();
  },

  computed: {
    getRoom() {
      return this.$store.state.roomid;
    }
  },

  methods: {
    loadData() {
      this.axios
        .get(`${this.$hostname}/Game/api/TicTacToeBox/Game/wait`, {
          withCredentials: true,
          params: { room: this.getRoom }
        })
        .then(response => {
          this.collection1 = response.data.BoxCollection1;
          this.collection2 = response.data.BoxCollection2;
          this.board = response.data.Board;
          this.boardPlayer = response.data.BoardPlayer;
          this.thisPlayer = response.data.ThisPlayer;
          this.player1 = response.data.Player1;
          this.player2 = response.data.Player2;
          this.round = response.data.Round;
          this.playerNow = response.data.PlayerNow;

          if (this.collection1.length < 6 || this.collection2.length < 6) {
            this.updateBoardDisplay();
          }
        })
        .catch(err => {
          this.errors.push(err);
        });

      if (this.player2 == "") {
        window.setTimeout(this.loadData, 1000);
      } else {
        this.queryServer();
      }
    },

    queryServer() {
      this.axios
        .get(`${this.$hostname}/Game/api/TicTacToeBox/Game/room`, {
          withCredentials: true,
          params: {
            room: this.getRoom
          }
        })
        .then(response => {
          this.collection1 = response.data.BoxCollection1;
          this.collection2 = response.data.BoxCollection2;
          this.board = response.data.Board;
          this.boardPlayer = response.data.BoardPlayer;
          this.round = response.data.Round;
          if (this.playerNow != response.data.PlayerNow) {
            this.updateBoardDisplay();
          }
          this.playerNow = response.data.PlayerNow;
          this.winner = response.data.Winner;
        });
      if (this.winner == "") {
        window.setTimeout(this.queryServer, 1000);
      }
    },

    updateBoardDisplay() {
      for (let i = 0; i < 3; i++) {
        for (let j = 0; j < 3; j++) {
          var len = this.board[i][j].length;
          if (len == 0) {
            this.boardDisplay[i][j] = "";
            this.boardDisplayPlayer[i][j] = "";
          } else {
            this.boardDisplay[i][j] = this.board[i][j][len - 1];
            this.boardDisplayPlayer[i][j] = this.boardPlayer[i][j][len - 1];
          }
        }
      }
    },

    chooseBoxFromBoard(i, j) {
      if (
        this.boardDisplayPlayer[i][j] == this.playerNow &&
        this.playerNow == this.thisPlayer
      ) {
        this.notice = "";
        this.boxChosenIndex = -1;
        this.boxChosenI = i;
        this.boxChosenJ = j;
        this.boxChosen = this.boardDisplay[i][j];
      }
    },

    choosebox(index) {
      if (this.playerNow === "player1" && this.thisPlayer === "player1") {
        this.notice = "";
        this.boxChosenI = -1;
        this.boxChosenJ = -1;
        this.boxChosenIndex = index;
        this.boxChosen = this.collection1[index];
        return;
      }
      if (this.playerNow === "player2" && this.thisPlayer === "player2") {
        this.notice = "";
        this.boxChosenI = -1;
        this.boxChosenJ = -1;
        this.boxChosenIndex = index;
        this.boxChosen = this.collection2[index];
      }
    },

    chooseCell(i, j) {
      this.notice = "";

      // Check if they choose a box, if yes, then do the rest
      if (this.boxChosenIndex != -1) {
        // Check if the player could place it there
        var avail = this.checkPlaceAvailble(i, j);

        if (avail) {
          this.board[i][j].push(this.boxChosen);
          this.boardPlayer[i][j].push(this.thisPlayer);
          this.boardDisplay[i][j] = this.boxChosen;
          this.boardDisplayPlayer[i][j] = this.thisPlayer;

          if (this.thisPlayer == "player1") {
            this.collection1.splice(this.boxChosenIndex, 1);
            this.axios
              .post(
                `${this.$hostname}/Game/api/TicTacToeBox/Game/room`,
                {
                  BoxCollection1: this.collection1,
                  Board: this.board,
                  BoardPlayer: this.boardPlayer,
                  BoardShowPlayer: this.boardDisplayPlayer
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
          } else {
            this.collection2.splice(this.boxChosenIndex, 1);
            this.axios
              .post(
                `${this.$hostname}/Game/api/TicTacToeBox/Game/room`,
                {
                  BoxCollection2: this.collection2,
                  Board: this.board,
                  BoardPlayer: this.boardPlayer,
                  BoardShowPlayer: this.boardDisplayPlayer
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
          }
          this.boxChosenIndex = -1;
          this.boxChosen = "";
        } else {
          this.notice = "Sorry! You can't place the box there!";
        }
      } else {
        // this.boxChosenIndex == -1
        // Which means the user may choose from board
        if (this.boxChosenI != -1 && this.boxChosenJ != -1) {
          avail = this.checkPlaceAvailble(i, j);

          if (avail) {
            this.board[this.boxChosenI][this.boxChosenJ].pop();
            this.boardPlayer[this.boxChosenI][this.boxChosenJ].pop();
            this.board[i][j].push(this.boxChosen);
            this.boardPlayer[i][j].push(this.thisPlayer);
            this.updateBoardDisplay();

            this.axios
              .post(
                `${this.$hostname}/Game/api/TicTacToeBox/Game/room`,
                {
                  BoxCollection1: this.collection1,
                  BoxCollection2: this.collection2,
                  Board: this.board,
                  BoardPlayer: this.boardPlayer,
                  BoardShowPlayer: this.boardDisplayPlayer
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

            this.boxChosen = "";
            this.boxChosenI = -1;
            this.boxChosenJ = -1;
          } else {
            this.notice = "Sorry! You can't place the box there!";
          }
        }
      }
    },

    checkPlaceAvailble(i, j) {
      var boxNow = this.boardDisplay[i][j];
      switch (boxNow) {
        case "":
          return true;
        case "small":
          if (this.boxChosen == "medium" || this.boxChosen == "large") {
            return true;
          }
          return false;
        case "medium":
          if (this.boxChosen == "large") {
            return true;
          }
          return false;
        default:
          return false;
      }
    }
  }
};
</script>

<style scoped>
.game-round-player {
  width: 100%;
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
.notice {
  background-color: #c6e2ff;
  background-image: linear-gradient(to right, #c6e2ff, white);
  margin-bottom: 40px;
}

.competition {
  display: grid;
  grid-template-columns: 110px auto 110px;
  grid-gap: 20px;
}
.competition__board {
  text-align: center;
}
.cube__small {
  font-size: 2em;
}
.cube__medium {
  font-size: 4em;
}
.cube__large {
  font-size: 6em;
}
.blue {
  color: #00009a;
}
.green {
  color: #1b5f1b;
}
.chosen {
  background-color: rgba(0, 0, 0, 0.3);
}

table {
  border: 1px solid black;
  border-collapse: collapse;
}
td {
  width: 120px;
  height: 120px;
  border: 1px solid black;
  border-collapse: collapse;
}
td div {
  width: 100%;
  height: 100%;
}
</style>
