<template>
  <div class="introduction-page">
    <h1>
      <span class="fa fa-rotate-right"></span>
      <router-link :to="{name: 'Home'}" class="link__none-style">Rotating puzzle</router-link>
    </h1>

    <div class="introduction-describe">
      <h2>introduction:</h2>
      <div>
        The four color blocks of red, yellow, blue, and green are
        distributed out of order on the square board.
        There is a rotation button in the middle of each color block.
        When the player clicks the button, the adjacent 4 color blocks will rotate 90 degrees clockwise.
        To win this game, the color blocks should be distributed the same
        as the target pattern by clicking the button to rotating the block.
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
      <button v-on:click="startLocalGame()" class="introduction-button">Play on this device!</button>
    </div>
  </div>
</template>

<script>
export default {
  data: function() {
    return {
      board: "easy"
    };
  },
  methods: {
    startLocalGame() {
      if (this.board == "hell") {
        this.$router.push({
          path: "/Game/RotatingPuzzle/Localgame/Hard"
        });
      } else {
        this.$router.push({
          path: "/Game/RotatingPuzzle/Localgame",
          query: { board: this.board }
        });
      }
    }
  }
};
</script>


<style scoped>
select {
  font-family: "Neucha", sans-serif;
  font-size: 1em;
  width: min-content;
  padding: 2px 5px 0 5px;
}
option {
  font-family: "Neucha", sans-serif;
  font-size: 1em;
  width: min-content;
}

.introduction-describe {
  font-family: "Neucha", sans-serif;
  font-size: 1.25em;
}
h2 {
  font-family: "Aladin", cursive;
  font-weight: bold;
}
.introduction-button {
  font-family: "Neucha", sans-serif;
  font-size: 1.25em;
  padding: 4px 10px 0 10px;
  border-radius: 3px;
  margin-top: 20px;
  margin-left: 50%;
}
button:hover {
  background-color: #003bac;
  border-top: 2px solid #608cdf;
  border-left: 2px solid #608cdf;
  border-bottom: 2px solid #002a7b;
  border-right: 2px solid #002a7b;
  color: white;
}

@media (max-width: 700px) {
  .introduction-describe {
    font-size: 1em;
  }
  .introduction-button {
    font-size: 1em;
  }
}
</style>
