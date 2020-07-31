<template>
  <div class="page">
    <h1>
      <span class="fa fa-rotate-right"></span>
      <router-link :to="{name: 'Home'}" class="link__none-style">Rotating puzzle</router-link>
    </h1>

    <div class="game-round-player">
      <div class="game-congratulation">{{win}}</div>
      <div class="board">
        <div v-for="(line, indexline) in board" :key="indexline">
          <div
            v-for="(cell, indexcell) in line"
            :key="indexcell"
            class="triangle"
            v-bind:class="{'rotate': indexline % 2 == 0 && indexcell % 2 == 1 ||
        indexline % 2 == 1 && indexcell % 2 == 0,
        'light-green': cell == 'light-green',
        'yellow':cell == 'yellow',
        'green': cell == 'green',
        'orange': cell == 'orange',
        'purple': cell == 'purple',
        'pink': cell == 'pink'}"
          ></div>
          <br>
        </div>

        <div class="button-board">
          <div
            v-for="(line,indexline) in buttonArray"
            :key="indexline"
            class="button-line"
            v-bind:class="{'button-move-left': indexline % 2 == 0}"
          >
            <div
              v-for="(cell, indexcell) in line"
              :key="indexcell"
              class="button-cell"
              v-bind:class="{'button-last':indexcell == 5}"
            >
              <button
                v-if="cell == 1"
                class="fa fa-rotate-right"
                v-on:click="rotate(indexline, indexcell)"
              ></button>
              <button v-else class="button-empty"></button>
            </div>
          </div>
        </div>
      </div>

      <div class="center">Target:</div>
      <div>
        <div v-for="(line, indexline) in boardExample" :key="indexline">
          <div
            v-for="(cell, indexcell) in line"
            :key="indexcell"
            class="triangle_example"
            v-bind:class="{'rotate': indexline % 2 == 0 && indexcell % 2 == 1 ||
        indexline % 2 == 1 && indexcell % 2 == 0,
        'light-green-example': cell == 'light-green',
        'yellow-example':cell == 'yellow',
        'green-example': cell == 'green',
        'orange-example': cell == 'orange',
        'purple-example': cell == 'purple',
        'pink-example': cell == 'pink'}"
          ></div>
          <br>
        </div>
      </div>

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
  data: function() {
    return {
      board: [],
      boardExample: [],
      boardExampleColor: [],
      panel: [9, 9, 9, 9, 9, 9],
      buttonArray: [],
      win: ""
    };
  },
  mounted: function() {
    // Initate board
    var line = [];
    for (let i = 0; i < 11; i++) {
      line[i] = "";
    }

    for (let i = 0; i < 6; i++) {
      this.board.push(line.slice());
    }

    // generate color for board
    for (let i = 0; i < 6; i++) {
      for (let j = 0; j < 11; j++) {
        if (
          ((i == 0 || i == 5) && (j == 0 || j == 1 || j == 9 || j == 10)) ||
          ((i == 1 || i == 4) && (j == 0 || j == 10))
        ) {
          continue;
        }
        this.board[i][j] = this.generateColor();
      }
    }

    // generate button
    for (let i = 0; i < 5; i++) {
      this.buttonArray.push([]);
      for (let j = 0; j < 5; j++) {
        if (
          ((i == 0 || i == 4) && (j == 0 || j == 4)) ||
          ((i == 1 || i == 3) && j == 0)
        ) {
          this.buttonArray[i].push(0);
          continue;
        }
        this.buttonArray[i].push(1);
      }
    }

    // generate color for example
    this.panel = [1, 1, 1, 1, 1, 1];
    for (let i = 0; i < 6; i++) {
      this.boardExampleColor.push(this.generateColor());
    }
    this.boardExample.push([
      "",
      "",
      this.boardExampleColor[0],
      this.boardExampleColor[1],
      this.boardExampleColor[1],
      this.boardExampleColor[1],
      this.boardExampleColor[1],
      this.boardExampleColor[1],
      this.boardExampleColor[2],
      "",
      ""
    ]);
    this.boardExample.push([
      "",
      this.boardExampleColor[0],
      this.boardExampleColor[0],
      this.boardExampleColor[0],
      this.boardExampleColor[1],
      this.boardExampleColor[1],
      this.boardExampleColor[1],
      this.boardExampleColor[2],
      this.boardExampleColor[2],
      this.boardExampleColor[2],
      ""
    ]);
    this.boardExample.push([
      this.boardExampleColor[0],
      this.boardExampleColor[0],
      this.boardExampleColor[0],
      this.boardExampleColor[0],
      this.boardExampleColor[0],
      this.boardExampleColor[1],
      this.boardExampleColor[2],
      this.boardExampleColor[2],
      this.boardExampleColor[2],
      this.boardExampleColor[2],
      this.boardExampleColor[2]
    ]);
    this.boardExample.push([
      this.boardExampleColor[3],
      this.boardExampleColor[3],
      this.boardExampleColor[3],
      this.boardExampleColor[3],
      this.boardExampleColor[3],
      this.boardExampleColor[4],
      this.boardExampleColor[5],
      this.boardExampleColor[5],
      this.boardExampleColor[5],
      this.boardExampleColor[5],
      this.boardExampleColor[5]
    ]);
    this.boardExample.push([
      "",
      this.boardExampleColor[3],
      this.boardExampleColor[3],
      this.boardExampleColor[3],
      this.boardExampleColor[4],
      this.boardExampleColor[4],
      this.boardExampleColor[4],
      this.boardExampleColor[5],
      this.boardExampleColor[5],
      this.boardExampleColor[5],
      ""
    ]);
    this.boardExample.push([
      "",
      "",
      this.boardExampleColor[3],
      this.boardExampleColor[4],
      this.boardExampleColor[4],
      this.boardExampleColor[4],
      this.boardExampleColor[4],
      this.boardExampleColor[4],
      this.boardExampleColor[5],
      "",
      ""
    ]);
  },

  methods: {
    generateColor() {
      for (let i = 0; i < 20; i++) {
        var r = Math.floor(Math.random() * 6);
        switch (r) {
          case 0:
            if (this.panel[r] > 0) {
              this.panel[r] = this.panel[r] - 1;
              return "light-green";
            }
            break;
          case 1:
            if (this.panel[r] > 0) {
              this.panel[r] = this.panel[r] - 1;
              return "yellow";
            }

            break;
          case 2:
            if (this.panel[r] > 0) {
              this.panel[r] = this.panel[r] - 1;
              return "green";
            }
            break;
          case 3:
            if (this.panel[r] > 0) {
              this.panel[r] = this.panel[r] - 1;
              return "orange";
            }
            break;
          case 4:
            if (this.panel[r] > 0) {
              this.panel[r] = this.panel[r] - 1;
              return "purple";
            }
            break;
          case 5:
            if (this.panel[r] > 0) {
              this.panel[r] = this.panel[r] - 1;
              return "pink";
            }
            break;

          default:
            break;
        }
      }
    },

    rotate(row, column) {
      switch (row) {
        case 0:
          var startColumn = column * 2;
          break;
        case 1:
          startColumn = column * 2 - 1;
          break;
        case 2:
          startColumn = column * 2;
          break;
        case 3:
          startColumn = column * 2 - 1;
          break;
        case 4:
          startColumn = column * 2;
          break;
        default:
          break;
      }
      var startRow = row;
      var startItem = this.board[startRow][startColumn];

      // Change color!
      this.$set(
        this.board[startRow],
        startColumn,
        this.board[startRow + 1][startColumn]
      );
      this.$set(
        this.board[startRow + 1],
        startColumn,
        this.board[startRow + 1][startColumn + 1]
      );
      this.$set(
        this.board[startRow + 1],
        startColumn + 1,
        this.board[startRow + 1][startColumn + 2]
      );
      this.$set(
        this.board[startRow + 1],
        startColumn + 2,
        this.board[startRow][startColumn + 2]
      );
      this.$set(
        this.board[startRow],
        startColumn + 2,
        this.board[startRow][startColumn + 1]
      );
      this.$set(this.board[startRow], startColumn + 1, startItem);

      // Check winner
      let win = this.checkSuccess();
      if (win) {
        this.win = "Congratulation!";
      }
    },

    checkSuccess() {
      for (let i = 0; i < 6; i++) {
        for (let j = 0; j < 11; j++) {
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
* {
  box-sizing: border-box;
}
.triangle {
  display: inline-block;
  margin-right: -25px;
  width: 0;
  height: 0;
  border-left: 30px solid transparent;
  border-right: 30px solid transparent;
  border-bottom: 60px solid white;
}
.rotate {
  transform: rotate(180deg);
}
.light-green {
  border-bottom: 60px solid #96ceb4;
}
.yellow {
  border-bottom: 60px solid #ffeead;
}
.green {
  border-bottom: 60px solid #3b9a9c;
}
.orange {
  border-bottom: 60px solid #ffad60;
}
.purple {
  border-bottom: 60px solid #a696c8;
}
.pink {
  border-bottom: 60px solid #fad3cf;
}

.triangle_example {
  display: inline-block;
  margin-top: 4px;
  margin-right: -10px;
  width: 0;
  height: 0;
  border-left: 15px solid transparent;
  border-right: 15px solid transparent;
  border-bottom: 15px solid white;
  vertical-align: top;
}
.light-green-example {
  border-bottom: 30px solid #96ceb4;
}
.yellow-example {
  border-bottom: 30px solid #ffeead;
}
.green-example {
  border-bottom: 30px solid #3b9a9c;
}
.orange-example {
  border-bottom: 30px solid #ffad60;
}
.purple-example {
  border-bottom: 30px solid #a696c8;
}
.pink-example {
  border-bottom: 30px solid #fad3cf;
}

.board {
  position: relative;
  width: max-content;
  margin-left: auto;
  margin-right: auto;
  margin-top: 15px;
  margin-bottom: 15px;
}
.button-board {
  position: absolute;
  top: 18px;
  left: 10px;
}
.button-line {
  display: flex;
  margin-top: 23px;
}
.button-cell {
  display: inline-block;
  vertical-align: top;
  margin-right: 30px;
}
button {
  width: 40px;
  height: 40px;
  font-size: 20px;
  border-radius: 20px;
}
.button-last {
  margin-right: 0px;
}
.button-empty {
  border: none;
  background-color: transparent;
}
.button-move-left {
  margin-left: 35px;
}

.game-round-player {
  width: 100%;
  text-align: center;
  font-family: "Aladin", cursive;
  font-size: 1.25em;
  font-weight: bold;
  margin-bottom: 10px;
}
.game-congratulation {
  color: red;
}

.introduction-button {
  font-family: "Neucha", sans-serif;
  font-size: 1.2em;
  margin: 15px 0;
  padding: 4px 10px 0 10px;
  border-radius: 3px;
  color: black;
  width: max-content;
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

@media (max-width: 450px) {
  .triangle {
    margin-right: -15px;
    margin-top: -2px;
    border-left: 20px solid transparent;
    border-right: 20px solid transparent;
    border-bottom: 40px solid white;
  }

  .light-green {
    border-bottom: 40px solid #96ceb4;
  }
  .yellow {
    border-bottom: 40px solid #ffeead;
  }
  .green {
    border-bottom: 40px solid #3b9a9c;
  }
  .orange {
    border-bottom: 40px solid #ffad60;
  }
  .purple {
    border-bottom: 40px solid #a696c8;
  }
  .pink {
    border-bottom: 40px solid #fad3cf;
  }

  .button-board {
    top: 15px;
    left: 5px;
  }
  .button-line {
    margin-top: 10px;
  }
  .button-cell {
    margin-right: 20px;
  }
  button {
    width: 30px;
    height: 30px;
    font-size: 18px;
    border-radius: 15px;
  }
  .button-move-left {
    margin-left: 25px;
  }
}
</style>
