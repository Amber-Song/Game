<template>
  <div class="introduction-page">
    <h1>
      <router-link :to="{name: 'Home'}" class="link__none-style">Monopoly</router-link>
    </h1>

    <div class="introduction-describe">
      <h2>introduction:</h2>
      <div></div>
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
      this.$router.push({
        path: "/Game/Monopoly/Localgame",
        query: { playerNum: this.playerNum }
      });
    }
  }
};
</script>

