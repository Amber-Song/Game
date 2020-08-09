<template>
  <div class="introduction-page">
    <h1>
      <span class="fa fa-hand-scissors-o icon__scissor"></span>
      <router-link :to="{name: 'Home'}" class="link__none-style">Rock Paper Scissors</router-link>
    </h1>

    <div class="introduction-describe">
      <h2>introduction:</h2>
      <div>
        Each player gets 3 cards with rock, paper and scissor on it. Cards are random chosen so the three cards may be the same.
        The player could only see his own cards while playing.
        In each turn, the player has to decide which card he wants to play. After both players decide their card,
        the two cards will be flipped together. The loser of each turn takes two cards back to his card collection.
        If it is drawn, the cards go back to its own collection. If one player doesn't play any card, he loses in this round.
        The one who has no card in hand is the final winner.
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
        .get(`${this.$hostname}/Game/api/RockPaperScissor/Game`, {
          withCredentials: true
        })
        .then(response => {
          // this.$router.push({
          //   path: "/Game/RockPaperScissors/Game/room",
          //   query: { room: response.data }
          // });
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
}
button:hover {
  background-color: #003bac;
  border-top: 2px solid #608cdf;
  border-left: 2px solid #608cdf;
  border-bottom: 2px solid #002a7b;
  border-right: 2px solid #002a7b;
  color: white;
}
</style>
