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
          <div>
            <div v-for="(line, indexline) in boardExample" :key="indexline">
              <div
                v-for="(cell, indexcell) in line"
                :key="indexcell"
                class="triangle_example"
                v-bind:class="{'triangle_rotate': indexline % 2 == 0 && indexcell % 2 == 1 ||
        indexline % 2 == 1 && indexcell % 2 == 0,
        'triangle_example_light-green': cell == 'light-green',
        'triangle_example_light-yellow':cell == 'yellow',
        'triangle_example_green': cell == 'green',
        'triangle_example_orange': cell == 'orange',
        'triangle_example_purple': cell == 'purple',
        'triangle_example_pink': cell == 'pink'}"
              ></div>
              <br>
            </div>
          </div>
        </div>

        <div class="board">
          <div v-for="(line, indexline) in board" :key="indexline">
            <div
              v-for="(cell, indexcell) in line"
              :key="indexcell"
              class="triangle"
              v-bind:class="{'triangle_rotate': indexline % 2 == 0 && indexcell % 2 == 1 ||
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
                  class="fa fa-rotate-right button-rotate"
                  v-on:click="rotate(indexline, indexcell)"
                ></button>
                <button v-else class="button-empty button-rotate"></button>
              </div>
            </div>
          </div>
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
  data: function() {
    return {
      board: [],
      boardExample: [],
      boardExampleColor: [],
      panel: [9, 9, 9, 9, 9, 9],
      buttonArray: [],
      win: "",
      timer: 0,
      gameTimer: null
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

    // Set timer
    this.gameTimer = window.setTimeout(this.updateTimer, 1000);
  },

  beforeDestroy() {
    this.clearTimer();
  },

  methods: {
    generateColor() {
      for (let i = 0; i < 30; i++) {
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
      this.board[startRow][startColumn] = this.board[startRow + 1][startColumn];
      this.board[startRow + 1][startColumn] = this.board[startRow + 1][startColumn + 1];
      this.board[startRow + 1][startColumn + 1] = this.board[startRow + 1][startColumn + 2];
      this.board[startRow + 1][startColumn + 2] = this.board[startRow][startColumn + 2];
      this.board[startRow][startColumn + 2] = this.board[startRow][startColumn + 1];
      this.board[startRow][startColumn + 1] = startItem;

      let isWin = this.checkSuccess();
      if (isWin) {
        this.win = "Congratulation! You win!";
        this.clearTimer();
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
  grid-template-columns: 240px auto;
  grid-gap: 10px;
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
.button-rotate {
  width: 40px;
  height: 40px;
  font-size: 20px;
  border-radius: 20px;
}
.button-rotate:hover {
  background: #efefef;
  color: black;
  border: initial;
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

@media (max-width: 700px) {
  .competition-section {
    display: block;
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
  .button-rotate {
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
