<template>
  <div class="introduction-page">
    <h1>
      <router-link :to="{name: 'Home'}" class="link__none-style">Monopoly</router-link>
    </h1>

    <div class="introduction-describe">
      <h2>introduction:</h2>
      <div>
        Each player has 30 coins and 15 pieces of grass.
        <br>
        <br>When the game start, the player could choose the character to head forward, left or right, then walk several steps based on how many dice the player got. When the character hits the wall, it will turn following the arrow.
        <br>

        <br>After the character stops, if the character is standing on other player's grass, the player has to pay the other player coins. The payment is based on how big that area is.
        Each block is calculated as 1 coin and it only calculates the grass on the top. If the character stops on its own grass or empty ground, the player doesn't need to pay.
        <br>
        <br>After payment, the player can put a piece of grass next to the character. Each piece of grass can cover 2 blocks and both two blocks have to be next to the character. The player can only put grass on the board in these three situations:
        <br>
        <br>1. The two blocks are empty
        <br>2. One of the block is empty and the other one is covered with any player's grass
        <br>3. Two blocks are both covered with grass, but the grasses are from a different player.
        <br>
        <br>After all the players use up all the grass, the one who has the most coins wins. If several players have the same amount of coins, the one who covers more area win.
      </div>

      <label>
        Player Number:
        <input type="number" v-model="playerNum" min="2" max="4">
      </label>
      <!-- <button v-on:click="startGame()" class="introduction-button">Play Online!</button> -->
      <button v-on:click="startLocalGame()" class="introduction-button">Play on this device!</button>
    </div>
  </div>
</template>

<script>
export default {
  data: function() {
    return { playerNum: 2 };
  },
  methods: {
    // startGame() {
    //   this.axios
    //     .get(`${this.$hostname}/Game/api/TicTacToeBox/Game`, {
    //       withCredentials: true
    //     })
    //     .then(response => {
    //       alert(
    //         "Please copy the address shown on next page and send to the other player to invite she/he to enter the game!"
    //       );
    //       this.$router.push({
    //         path: "/Game/TicTacToeBox/Game/room",
    //         query: { room: response.data }
    //       });
    //     })
    //     .catch(error => {
    //       console.log("Error:", error); // Logs out the error
    //     });
    // },

    startLocalGame() {
      if (this.playerNum > 4) {
        this.playerNum = 4;
      }
      if (this.playerNum < 2) {
        this.playerNum = 2;
      }
      this.$router.push({
        path: "/Game/Monopoly/Localgame",
        query: { playerNum: this.playerNum }
      });
    }
  }
};
</script>

