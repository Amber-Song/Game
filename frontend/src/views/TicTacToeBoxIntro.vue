<template>
  <div class="introduction-page">
    <h1>
      <span class="fa fa-cube title__icon"></span>
      <router-link :to="{name: 'Home'}" class="link__none-style">Tic Tac Toe Box version</router-link>
    </h1>

    <div class="introduction-describe">
      <h2>introduction:</h2>
      <div>
        Each player has two small, medium and large size boxes.
        The player places or moves one box onto the 3 by 3 board, per turn.
        To move or place boxes (by clicking the box and board), you should follow the following rules:
        <ol>
          <li>Boxes placed or moved on the board can cover a box which is smaller than it or be placed on a blank area.</li>
          <li>The box a player wants to move has to be the top box; you cannot move a box covered by a larger box.</li>
          <li>If a box on the board is moved, any box it covered will appear.</li>
        </ol>A player will win the game if three of their boxes, of any size, are in a line, row or diagonal.
      </div>
      <button v-on:click="startGame()" class="introduction-button">Play Online!</button>
      <button v-on:click="startLocalGame()" class="introduction-button">Play on this device!</button>
    </div>
  </div>
</template>

<script>
export default {
  methods: {
    startGame() {
      this.axios
        .get(`${this.$hostname}/Game/api/TicTacToeBox/Game`, {
          withCredentials: true
        })
        .then(response => {
          alert(
            "Please copy the address shown on next page and send to the other player to invite she/he to enter the game!"
          );
          this.$router.push({
            path: "/TicTacToeBox/Game/room",
            query: { room: response.data }
          });
        })
        .catch(error => {
          console.log("Error:", error); // Logs out the error
        });
    },

    startLocalGame() {
      this.$router.push({
        path: "/TicTacToeBox/Localgame"
      });
    }
  }
};
</script>
