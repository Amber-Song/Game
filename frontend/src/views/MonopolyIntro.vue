<template>
  <div class="introduction-page">
    <h1>
      <img src="../assets/flower.png" style="height:30px;">
      <router-link :to="{name: 'Home'}" class="link__none-style">Monopoly</router-link>
    </h1>

    <div class="introduction-describe">
      <h2>introduction:</h2>
      <div>
        Each player has 30 coins and 15 pieces of grass.
        <br>
        <br>When the game start, the player could choose the character to head forward, left or right, then walk several steps based on how many dice the player got. When the character hits the wall, it will turn following the arrow.
        <br>
        <div class="example__direction-dice">
          <div>
            <img
              src="../assets/example_direction.png"
              alt="direction example"
              class="picture-width-240 pixelart"
            >
            <img
              src="../assets/example_dice.png"
              alt="dice example"
              class="picture-width-200 pixelart"
            >
          </div>
          <img src="../assets/example_walk.png" alt="walk example" class="picture-height-260">
        </div>

        <br>After the character stops, if the character is standing on other player's grass, the player has to pay the other player coins. The payment is based on how big that area is.
        Each block is calculated as 1 coin and it only calculates the grass on the top. If the character stops on its own grass or empty ground, the player doesn't need to pay.
        <br>

        <div class="example__direction-dice">
          <img
            src="../assets/example_payment.png"
            alt="payment example"
            class="picture-width-240 pixelart"
          >
          <img src="../assets/example_stand.png" alt="position example" class="picture-height-260">
        </div>

        <br>After payment, the player can put a piece of grass next to the character. Each piece of grass can cover 2 blocks and both two blocks have to be next to the character. The player can only put grass on the board in these three situations:
        <br>
        <br>1. The two blocks are empty
        <br>2. One of the block is empty and the other one is covered with any player's grass
        <br>3. Two blocks are both covered with grass, but the grasses are from a different player.
        <br>

        <div class="example__choose-block">
          <img
            src="../assets/example_situation1.png"
            alt="situation1 example"
            class="picture-height-150"
          >
          <img
            src="../assets/example_situation2.png"
            alt="situation2 example"
            class="picture-height-150"
          >
          <img
            src="../assets/example_situation3.png"
            alt="situation3 example"
            class="picture-height-150"
          >
        </div>

        <br>After all the players use up all the grass, the one who has the most coins wins. If several players have the same amount of coins, the one who covers more area win.
      </div>
      <br>
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

<style scoped>
.example__direction-dice {
  margin: 20px auto;
  display: grid;
  grid-template-columns: 240px 260px;
  justify-content: space-around;
}
.example__choose-block {
  margin: 20px auto;
  display: grid;
  grid-template-columns: 150px 150px 150px;
  justify-content: space-around;
}
.picture-width-240 {
  width: 240px;
}
.picture-width-200 {
  width: 200px;
}
.picture-height-260 {
  height: 260px;
}
.picture-height-150 {
  height: 150px;
}
.pixelart {
  image-rendering: pixelated;
}
</style>
