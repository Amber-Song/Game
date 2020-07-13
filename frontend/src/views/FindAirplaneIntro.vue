<template>
  <div class="introduction-page">
    <h1 title="Get bored of playing this? Take a look at other games!">
      <font-awesome-icon :icon="['fas', 'plane']" class="icon__airplane"/>
      <router-link :to="{name: 'Home'}" class="link__none-style">Seek for airplane's head</router-link>
    </h1>
    <!-- This is the image for the example -->
    <div class="introduction-content">
      <!-- This is the introduction -->
      <div class="introduction-describe">
        <h2>Introduction:</h2>
        <div>
          There are two airplanes whose shapes are shown in the picture on the board.
          Each player has one board but the two airplanes face different directions and
          place in different positions. The two players are going to click one piece at each time in turn
          and the flipped pieces will show you it is an airplane's head, or a part of the body, or not airplanes.
          The one who uses the least pieces to find two airplanes' head (the dark blue piece) win the game.
        </div>
        <br>
        <h2>Setting:</h2>
        <label>
          The shape of airplane:
          <select name="shape">
            <option value="1">Airplane 1</option>
          </select>
        </label>
        <br>
        <label>
          Choose the boardlength (8-10):
          <input type="number" min="8" max="10" v-model="boardlength">
        </label>
        <button v-on:click="startGame()" class="introduction-button">Play Now!</button>
      </div>

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
    </div>
  </div>
</template>

<script>
export default {
  data: function() {
    return {
      boardlength: 8
    };
  },
  computed: {
    exampleArrays() {
      return this.$store.state.exampleMatrix;
    }
  },
  methods: {
    startGame() {
      this.axios
        .get(`${this.$hostname}/Game/api/FindAirplane/Game`, {
          withCredentials: true,
          params: { boardLength: this.boardlength }
        })
        .then(response => {
          this.$store.commit("getRoom", { roomid: response.data });
          this.$router.push({
            path: "/Game/FindAirplane/Game/room",
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
.introduction-content {
  display: grid;
  grid-template-columns: auto 260px;
}
.board-td {
  width: 30px;
  height: 30px;
  background-color: darkgray;
  border: 1px solid black;
  border-collapse: collapse;
}
.introduction-describe {
  font-family: "Neucha", sans-serif;
  font-size: 1.25em;
  margin: 30px;
  margin-top: 0px;
}
h2 {
  font-family: "Aladin", cursive;
  font-weight: bold;
  margin: 0px;
  margin-bottom: 10px;
  margin-top: 10px;
}
.introduction-button {
  font-family: "Neucha", sans-serif;
  font-size: 1.2em;
  margin: 40px;
  margin-left: 50%;
  padding: 4px 10px 0 10px;
  border-radius: 3px;
}
table {
  margin-top: 40px;
}
label {
  line-height: 200%;
}
input {
  font-family: "Neucha", sans-serif;
  font-size: 1em;
  width: 50px;
  padding: 6px 0 0 5px;
}
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
button:hover {
  background-color: #003bac;
  border-top: 2px solid #608cdf;
  border-left: 2px solid #608cdf;
  border-bottom: 2px solid #002a7b;
  border-right: 2px solid #002a7b;
  color: white;
}

@media (max-width: 750px) {
  .introduction-content {
    display: block;
  }
  .introduction-describe {
    font-family: "Neucha", sans-serif;
    font-size: 1em;
    margin: 30px;
    margin-top: 0px;
  }
  .introduction-button {
    font-family: "Neucha", sans-serif;
    font-size: 1em;
    margin: 40px;
    margin-left: 50%;
    padding: 4px 10px 0 10px;
    border-radius: 3px;
  }
}
</style>
