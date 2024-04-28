<template>
  <div class="introduction-page">
    <h1>
      <span class="fa fa-rotate-right title__icon"></span>
      <router-link :to="{name: 'Home'}" class="link__none-style">Rotating puzzle</router-link>
    </h1>

    <div class="introduction-describe">
      <h2>introduction:</h2>
      <div>
        The four-colour blocks of red, yellow, blue, and green or six-colour blocks are distributed out of order on the square board.
        There are rotation buttons in the middle of each colour block. When the player clicks the button,
        the adjacent four or six colour blocks will rotate 90 degrees clockwise.
        To win this game, the colour blocks should be distributed the same as the target pattern by clicking the button to rotating the block.
      </div>
      <br>

      <label>
        Choose level:
        <select v-model="board">
          <option value="easy">easy</option>
          <option value="hard">hard</option>
          <option value="hell">hell</option>
        </select>
      </label>

      <div>Example:</div>
      <table v-if="this.board == 'easy'" class="example-display">
        <tr v-for="(line, indexline) in board4" :key="indexline">
          <td
            v-for="(cell, indexcell) in line"
            :key="indexcell"
            class="rotate-puzzle__example-td"
            v-bind:class="{'rotate-puzzle__red': cell == 'red', 'rotate-puzzle__yellow':cell == 'yellow', 'rotate-puzzle__green': cell == 'green', 'rotate-puzzle__blue':cell == 'blue'}"
          ></td>
        </tr>
      </table>

      <table v-else-if="this.board == 'hard'" class="example-display">
        <tr v-for="(line, indexline) in board6" :key="indexline">
          <td
            v-for="(cell, indexcell) in line"
            :key="indexcell"
            class="rotate-puzzle__example-td"
            v-bind:class="{'rotate-puzzle__red': cell == 'red', 'rotate-puzzle__yellow':cell == 'yellow', 'rotate-puzzle__green': cell == 'green', 'rotate-puzzle__blue':cell == 'blue'}"
          ></td>
        </tr>
      </table>

      <div v-else class="example-display">
        <div v-for="(line, indexline) in board11" :key="indexline">
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

      <span class="example-display arrow">&rarr;</span>

      <table v-if="this.board == 'easy'" class="example-display">
        <tr v-for="(line, indexline) in board4Target" :key="indexline">
          <td
            v-for="(cell, indexcell) in line"
            :key="indexcell"
            class="rotate-puzzle__example-td"
            v-bind:class="{'rotate-puzzle__red': cell == 'red', 'rotate-puzzle__yellow':cell == 'yellow', 'rotate-puzzle__green': cell == 'green', 'rotate-puzzle__blue':cell == 'blue'}"
          ></td>
        </tr>
      </table>

      <table v-else-if="this.board == 'hard'" class="example-display">
        <tr v-for="(line, indexline) in board6Target" :key="indexline">
          <td
            v-for="(cell, indexcell) in line"
            :key="indexcell"
            class="rotate-puzzle__example-td"
            v-bind:class="{'rotate-puzzle__red': cell == 'red', 'rotate-puzzle__yellow':cell == 'yellow', 'rotate-puzzle__green': cell == 'green', 'rotate-puzzle__blue':cell == 'blue'}"
          ></td>
        </tr>
      </table>

      <div v-else class="example-display">
        <div v-for="(line, indexline) in board11Target" :key="indexline">
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

      <button v-on:click="startLocalGame()" class="introduction-button">Play on this device!</button>
    </div>
  </div>
</template>

<script>
import { mapStores } from 'pinia';
import { usePuzzleStore } from '../stores/puzzle';

export default {
  data: function() {
    return {
      board: "easy"
    };
  },
  computed: {
    ...mapStores(usePuzzleStore),
    board4() {
      return this.puzzleStore.rotatePuzzle4Example;
    },
    board4Target() {
      return this.puzzleStore.rotatePuzzle4Target;
    },
    board6() {
      return this.puzzleStore.rotatePuzzle6Example;
    },
    board6Target() {
      return this.puzzleStore.rotatePuzzle6Target;
    },
    board11() {
      return this.puzzleStore.rotatePuzzle11Example;
    },
    board11Target() {
      return this.puzzleStore.rotatePuzzle11Target;
    }
  },
  methods: {
    startLocalGame() {
      if (this.board == "hell") {
        this.$router.push({
          path: "/RotatingPuzzle/Localgame/Hard"
        });
      } else {
        this.$router.push({
          path: "/RotatingPuzzle/Localgame",
          query: { board: this.board }
        });
      }
    }
  }
};
</script>

<style scoped>
.example-display {
  display: inline-block;
  vertical-align: middle;
  margin-bottom: 10px;
}
.arrow {
  font-size: 2em;
  margin: 20px;
  margin-left: 30px;
}
</style>
