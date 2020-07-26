<template>
  <div class="page">
    <h1>
      <span class="fa fa-rotate-right"></span>
      <router-link :to="{name: 'Home'}" class="link__none-style">Rotating puzzle</router-link>
    </h1>

    <div class="game-round-player">
      <div class="game-congratulation">{{win}}</div>
      <div class="board">
        <table>
          <tr v-for="(line,lineIndex) in board" :key="lineIndex">
            <td v-for="(cell, cellIndex) in line" :key="cellIndex">
              <div v-if="cell == 'red'" class="cell__red"></div>
              <div v-else-if="cell == 'yellow'" class="cell__yellow"></div>
              <div v-else-if="cell == 'green'" class="cell__green"></div>
              <div v-else class="cell__blue"></div>
            </td>
          </tr>
        </table>
        <table class="board-button">
          <tr v-for="i in buttonNum" :key="i">
            <td v-for="j in buttonNum" :key="j" class="board-td">
              <button v-on:click="rotate(i,j)" class="fa fa-rotate-right button-rotate"></button>
            </td>
          </tr>
        </table>
      </div>

      <div class="center">Target:</div>
      <table>
        <tr v-for="(line, lineIndex) in boardExample" :key="lineIndex">
          <td v-for="(cell, cellIndex) in line" :key="cellIndex" class="example">
            <div v-if="cell == 'red'" class="cell__red"></div>
            <div v-else-if="cell == 'yellow'" class="cell__yellow"></div>
            <div v-else-if="cell == 'green'" class="cell__green"></div>
            <div v-else class="cell__blue"></div>
          </td>
        </tr>
      </table>
      <br>

      <button class="introduction-button">
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
      win: ""
    };
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
  },

  methods: {
    // This return a random color
    getRandomColor() {
      for (let i = 0; i < 20; i++) {
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
      this.$set(this.board[i], j, this.board[i - 1][j]);
      this.$set(this.board[i - 1], j, this.board[i - 1][j - 1]);
      this.$set(this.board[i - 1], j - 1, this.board[i][j - 1]);
      this.$set(this.board[i], j - 1, color);

      let win = this.checkSuccess();
      if (win) {
        this.win = "Congratulation!";
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
    }
  }
};
</script>

<style scoped>
.board {
  position: relative;
  width: max-content;
  margin-left: auto;
  margin-right: auto;
  margin-top: 15px;
  margin-bottom: 15px;
}
.board-button {
  position: absolute;
  top: 50px;
  left: 22px;
  text-align: end;
}
.board-td {
  border: none;
}

td {
  width: 100px;
  height: 100px;
  border: 1px solid black;
  border-collapse: collapse;
}
.cell__red {
  width: 100%;
  height: 100%;
  background-color: #ff3e3e;
}
.cell__yellow {
  width: 100%;
  height: 100%;
  background-color: #fafd51;
}
.cell__green {
  width: 100%;
  height: 100%;
  background-color: #77ec77;
}
.cell__blue {
  width: 100%;
  height: 100%;
  background-color: #4f4fff;
}

.button-rotate {
  height: 40px;
  width: 40px;
  font-size: 25px;
}

.example {
  width: 30px;
  height: 30px;
  border: 1px solid black;
  border-collapse: collapse;
}
.center {
  text-align: center;
}
.game-congratulation {
  color: red;
}
.game-round-player {
  width: 100%;
  text-align: center;
  font-family: "Aladin", cursive;
  font-size: 1.25em;
  font-weight: bold;
  margin-bottom: 10px;
}

.introduction-button {
  font-family: "Neucha", sans-serif;
  font-size: 1.2em;
  margin: 15px 0;
  padding: 4px 10px 0 10px;
  border-radius: 3px;
  color: black;
}
.introduction-button:hover {
  background-color: #003bac;
  border-top: 2px solid #608cdf;
  border-left: 2px solid #608cdf;
  border-bottom: 2px solid #002a7b;
  border-right: 2px solid #002a7b;
}
.introduction-button:hover a {
  color: white;
}

@media (max-width: 700px) {
  .introduction-button {
    font-size: 1em;
  }
}
@media (max-width: 440px) {
  .game-round-player {
    font-size: 1em;
  }
  .example {
    width: 20px;
    height: 20px;
  }
  .button-rotate {
    height: 30px;
    width: 30px;
    font-size: 20px;
  }
  .board-button {
    top: 25px;
    left: 22px;
  }
  td {
    width: 55px;
    height: 55px;
  }
}
</style>
