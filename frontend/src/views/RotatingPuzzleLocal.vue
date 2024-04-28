<template>
  <div class="game-page">
    <h1>
      <span class="fa fa-rotate-right title__icon"></span>
      <router-link :to="{name: 'Home'}" class="link__none-style">Rotating puzzle</router-link>
    </h1>

    <div class="game-content">
      <div class="game-congratulation">{{ win }}</div>

      <div class="competition-section">
        <div>
          <div>Time:</div>
          <div>{{ timer }} s</div>
          <br>
          <div>Target:</div>
          <table>
            <tr v-for="(line, indexline) in boardExample" :key="indexline">
              <td
                v-for="(cell, indexcell) in line"
                :key="indexcell"
                class="rotate-puzzle__example-td"
                v-bind:class="{'rotate-puzzle__red': cell == 'red', 'rotate-puzzle__yellow':cell == 'yellow', 'rotate-puzzle__green': cell == 'green', 'rotate-puzzle__blue':cell == 'blue'}"
              ></td>
            </tr>
          </table>
        </div>

        <div class="board">
          <table>
            <tr v-for="(line,indexline) in board" :key="indexline">
              <td
                v-for="(cell, indexcell) in line"
                :key="indexcell"
                class="board-td"
                v-bind:class="{'rotate-puzzle__red': cell == 'red', 'rotate-puzzle__yellow':cell == 'yellow', 'rotate-puzzle__green': cell == 'green', 'rotate-puzzle__blue':cell == 'blue'}"
              ></td>
            </tr>
          </table>
          <table class="board-button">
            <tr v-for="i in buttonNum" :key="i">
              <td v-for="j in buttonNum" :key="j" class="board-button-td board-td">
                <button v-on:click="rotate(i,j)" class="fa fa-rotate-right button-rotate"></button>
              </td>
            </tr>
          </table>
        </div>
      </div>

      <button class="introduction-button" v-on:click="clearTimer()">
        <router-link :to="{name: 'RotatingPuzzleIntroduction'}" class="link__none-style">
          Back to introduction page to
          <strong>restart</strong>.
        </router-link>
      </button>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    boardtype: String
  },
  data: function() {
    return {
      boardlength: 0,
      buttonNum: 0,
      board: [],
      panel: [],
      boardExample: [],
      boardExampleColor: [],
      win: "",
      timer: 0,
      gameTimer: null
    };
  },
  beforeDestroy() {
    this.clearTimer();
  },
  mounted: function() {
    switch (this.boardtype) {
      case "easy":
        this.boardlength = 4;
        this.panel = [4, 4, 4, 4];
        break;
      case "hard":
        this.boardlength = 6;
        this.panel = [9, 9, 9, 9];
        break;
      default:
        break;
    }

    // Make board
    this.buttonNum = this.boardlength - 1;
    for (let i = 0; i < this.boardlength; i++) {
      var line = [];
      for (let j = 0; j < this.boardlength; j++) {
        line.push(this.getRandomColor());
      }
      this.board.push(line);
    }

    // Make example
    this.panel = [1, 1, 1, 1];
    for (let i = 0; i < 4; i++) {
      this.boardExampleColor.push(this.getRandomColor());
    }

    if (this.boardtype == "easy") {
      this.boardExample.push([
        this.boardExampleColor[0],
        this.boardExampleColor[0],
        this.boardExampleColor[1],
        this.boardExampleColor[1]
      ]);
      this.boardExample.push([
        this.boardExampleColor[0],
        this.boardExampleColor[0],
        this.boardExampleColor[1],
        this.boardExampleColor[1]
      ]);
      this.boardExample.push([
        this.boardExampleColor[2],
        this.boardExampleColor[2],
        this.boardExampleColor[3],
        this.boardExampleColor[3]
      ]);
      this.boardExample.push([
        this.boardExampleColor[2],
        this.boardExampleColor[2],
        this.boardExampleColor[3],
        this.boardExampleColor[3]
      ]);
    }

    if (this.boardtype == "hard") {
      this.boardExample.push([
        this.boardExampleColor[0],
        this.boardExampleColor[0],
        this.boardExampleColor[0],
        this.boardExampleColor[1],
        this.boardExampleColor[1],
        this.boardExampleColor[1]
      ]);
      this.boardExample.push([
        this.boardExampleColor[0],
        this.boardExampleColor[0],
        this.boardExampleColor[0],
        this.boardExampleColor[1],
        this.boardExampleColor[1],
        this.boardExampleColor[1]
      ]);
      this.boardExample.push([
        this.boardExampleColor[0],
        this.boardExampleColor[0],
        this.boardExampleColor[0],
        this.boardExampleColor[1],
        this.boardExampleColor[1],
        this.boardExampleColor[1]
      ]);
      this.boardExample.push([
        this.boardExampleColor[2],
        this.boardExampleColor[2],
        this.boardExampleColor[2],
        this.boardExampleColor[3],
        this.boardExampleColor[3],
        this.boardExampleColor[3]
      ]);
      this.boardExample.push([
        this.boardExampleColor[2],
        this.boardExampleColor[2],
        this.boardExampleColor[2],
        this.boardExampleColor[3],
        this.boardExampleColor[3],
        this.boardExampleColor[3]
      ]);
      this.boardExample.push([
        this.boardExampleColor[2],
        this.boardExampleColor[2],
        this.boardExampleColor[2],
        this.boardExampleColor[3],
        this.boardExampleColor[3],
        this.boardExampleColor[3]
      ]);
    }

    // Set timer
    this.gameTimer = window.setTimeout(this.updateTimer, 1000);
  },

  methods: {
    // This return a random color
    getRandomColor() {
      for (let i = 0; i < 30; i++) {
        var r = Math.floor(Math.random() * 4);
        switch (r) {
          case 0:
            if (this.panel[0] > 0) {
              this.panel[0] = this.panel[0] - 1;
              return "red";
            }
            break;
          case 1:
            if (this.panel[1] > 0) {
              this.panel[1] = this.panel[1] - 1;
              return "yellow";
            }
            break;
          case 2:
            if (this.panel[2] > 0) {
              this.panel[2] = this.panel[2] - 1;
              return "green";
            }
            break;
          case 3:
            if (this.panel[3] > 0) {
              this.panel[3] = this.panel[3] - 1;
              return "blue";
            }
            break;
          default:
            break;
        }
      }
    },

    // This make the four color rotate
    rotate(i, j) {
      var color = this.board[i][j];
      this.board[i][j] = this.board[i - 1][j];
      this.board[i - 1][j] = this.board[i - 1][j - 1];
      this.board[i - 1][j - 1] = this.board[i][j - 1];
      this.board[i][j - 1] = color;

      let isWin = this.checkSuccess();
      if (isWin) {
        this.win = "Congratulation! You win!";
        this.clearTimer();
      }
    },

    checkSuccess() {
      for (let i = 0; i < this.boardlength; i++) {
        for (let j = 0; j < this.boardlength; j++) {
          if (this.board[i][j] != this.boardExample[i][j]) {
            return false;
          }
        }
      }
      return true;
    },

    updateTimer() {
      this.timer = this.timer + 1;
      this.gameTimer = window.setTimeout(this.updateTimer, 1000);
    },

    clearTimer() {
      window.clearTimeout(this.gameTimer);
    }
  }
};
</script>

<style scoped>
.competition-section {
  display: grid;
  grid-template-columns: 180px auto;
  grid-gap: 10px;
}

.board {
  position: relative;
  width: max-content;
  margin-left: auto;
  margin-right: auto;
  margin-top: 15px;
  margin-bottom: 15px;
}
.board-td {
  width: 100px;
  height: 100px;
  border: 1px solid black;
  border-collapse: collapse;
}

.board-button {
  position: absolute;
  top: 50px;
  left: 22px;
  text-align: end;
}
.board-button-td {
  border: none;
}
.button-rotate {
  height: 40px;
  width: 40px;
  font-size: 25px;
}

@media (max-width: 550px) {
  .competition-section {
    display: block;
  }
}

@media (max-width: 440px) {
  .rotate-puzzle__example-td {
    width: 20px;
    height: 20px;
  }
  .board-td {
    width: 55px;
    height: 55px;
  }
  .board-button {
    top: 25px;
  }
  .button-rotate {
    height: 30px;
    width: 30px;
    font-size: 20px;
  }
}
</style>
