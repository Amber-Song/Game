<template>
  <div class="introduction-page">
    <h1>
      <span class="fa fa-cube"></span>
      <router-link :to="{name: 'Home'}" class="link__none-style">Tic Tac Toe Box version</router-link>
    </h1>

    <div class="introduction-describe">
      <h2>introduction:</h2>
      <div>
        Each player have small, medium, large, three sizes of boxes, each size have two.
        The player place or move one box on the 3 by 3 board in turn.
        To move or place the boxes, it should follow the following rules:
        <ol>
          <li>The box put or moved on the board could cover the box which is smaller than it or place on the blank area.</li>
          <li>The box player wants to move has to be the top one which means it shouldn't be covered with larger boxes.</li>
          <li>Once the box on the board has moved the box it covered appeared.</li>
        </ol>The player will win the game if his three boxes are in a line, row or diagonal.
      </div>
      <button v-on:click="startGame()" class="introduction-button">Play Now!</button>
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
          this.$store.commit("getRoom", { roomid: response.data });
          this.$router.push({
            path: "/Game/TicTacToeBox/Game/room",
            query: { room: response.data }
          });
        })
        .catch(error => {
          console.log("Error:", error); // Logs out the error
        });
    }
  }
};
</script>


<style scoped>
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
