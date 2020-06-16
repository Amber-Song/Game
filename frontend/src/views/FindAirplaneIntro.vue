<template>
  <div class="introduction">
    <h1>&rarr; Find airplane's head</h1>
    <!-- This is the image for the example -->
    <div class="introduction-content">
      <table>
        <tr v-for="(array, indexarray) in exampleArrays" :key="indexarray">
          <td v-for="(block, indexblock) in array" :key="indexblock" class="board-td">
            <div v-if="block === 3" class="board-td__white"></div>
            <div v-else-if="block === 4" class="board-td__blue"></div>
            <div v-else-if="block === 5" class="board-td__darkblue"></div>
            <div v-else class="board-td__gray"></div>
          </td>
        </tr>
      </table>

      <!-- This is the introduction -->
      <div class="introduction-describe">
        <h2>Introduction:</h2>
        <div>
          There will be two airplanes shown in the picture on each board but
          facing the different directions and placing in a different position.
          The two players are going to click one piece at one time in turn. The
          one who uses the least pieces to find two airplanes head (the dark
          blue piece) win the game.
        </div>
        <button v-on:click="startGame()" class="introduction-button">Play Now!</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    exampleArrays() {
      return this.$store.state.exampleMatrix;
    }
  },
  methods: {
    startGame() {
      this.axios
        .get(`${this.$hostname}/FindAirplane/Game`, {
          withCredentials: true
        })
        .then(response => {
          this.$store.commit("getRoom", { roomid: response.data });
          this.$router.push({
            path: "/FindAirplane/Game/room",
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
.introduction {
  width: 80%;
  margin-left: auto;
  margin-right: auto;
}
.introduction-content {
  display: grid;
  grid-template-columns: 260px auto;
}
.board-td {
  width: 30px;
  height: 30px;
  background-color: darkgray;
  border: 1px solid black;
  border-collapse: collapse;
}
.introduction-describe {
  margin: 30px;
  margin-top: 10px;
}
h2 {
  font-weight: bold;
  margin: 0px;
  margin-bottom: 10px;
}
.introduction-button {
  margin: 20px;
  margin-left: 50%;
  padding: 5px 10px;
  border-radius: 3px;
}
</style>
