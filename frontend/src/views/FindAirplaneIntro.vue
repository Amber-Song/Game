<template>
  <div class="introduction-page">
    <h1>
      <font-awesome-icon :icon="['fas', 'plane']" class="icon__airplane title__icon"/>
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
          Play
          <select v-model="where">
            <option value="online">online</option>
            <option value="local">local</option>
          </select>
        </label>
        <br>
        <label>
          The shape of airplane:
          <select v-model="shape">
            <option value="airplaneA">Airplane A</option>
            <option value="airplaneB">Airplane B</option>
          </select>
        </label>
        <br>
        <label>
          Choose the boardlength (8-10):
          <input
            v-if="shape == 'airplaneA'"
            type="number"
            min="8"
            max="10"
            v-model="boardlength"
          >
          <input v-else type="number" min="10" max="10" v-model="boardlength">
        </label>
        <button v-on:click="startGame()" class="introduction-button">Play Now!</button>
      </div>

      <table v-if="shape == 'airplaneA'">
        <tr v-for="(array, indexarray) in airplaneA" :key="indexarray">
          <td v-for="(block, indexblock) in array" :key="indexblock" class="airplane-example__td">
            <div v-if="block === 3" class="airplane__white"></div>
            <div v-else-if="block === 4" class="airplane__blue"></div>
            <div v-else-if="block === 5" class="airplane__darkblue"></div>
            <div v-else class="airplane__gray"></div>
          </td>
        </tr>
      </table>

      <table v-else-if="shape == 'airplaneB'">
        <tr v-for="(array, indexarray) in airplaneB" :key="indexarray">
          <td v-for="(block, indexblock) in array" :key="indexblock" class="airplane-example__td">
            <div v-if="block === 3" class="airplane__white"></div>
            <div v-else-if="block === 4" class="airplane__blue"></div>
            <div v-else-if="block === 5" class="airplane__darkblue"></div>
            <div v-else class="airplane__gray"></div>
          </td>
        </tr>
      </table>
    </div>
  </div>
</template>

<script>
import { mapStores } from 'pinia';
import { useAirplaneStore } from '../stores/airplane';

export default {
  data: function() {
    return {
      boardlength: 10,
      where: "online",
      shape: "airplaneA"
    };
  },
  computed: {
    ...mapStores(useAirplaneStore),
    airplaneA() {
      return this.airplaneStore.airplaneA;
    },
    airplaneB() {
      return this.airplaneStore.airplaneB;
    }
  },
  methods: {
    startGame() {
      if (
        (this.boardlength > 10 || this.boardlength < 8) &&
        this.shape == "airplaneA"
      ) {
        this.boardlength = 8;
      }
      if (this.shape == "airplaneB") {
        this.boardlength = 10;
      }

      if (this.where == "local") {
        this.$router.push({
          path: "/Game/FindAirplane/Localgame",
          query: { boardlength: this.boardlength, shape: this.shape }
        });
      } else {
        this.axios
          .get(`${this.$hostname}/Game/api/FindAirplane/Game`, {
            withCredentials: true,
            params: { boardLength: this.boardlength, shape: this.shape }
          })
          .then(response => {
            // alert the room
            alert(
              "Please copy the address shown on next page and send to the other player to invite she/he to enter the game!"
            );
            this.$router.push({
              path: "/Game/FindAirplane/Game/room",
              query: { room: response.data, shape: this.shape }
            });
          })
          .catch(error => {
            console.log("Error:", error);
          });
      }
    }
  }
};
</script>

<style scoped>
.introduction-content {
  display: grid;
  grid-template-columns: auto 300px;
}

table {
  margin-top: 40px;
  width: max-content;
}
label {
  line-height: 200%;
}

@media (max-width: 700px) {
  .introduction-content {
    display: block;
  }
  .introduction-describe {
    font-size: 1em;
    margin: 10px;
    margin-top: 0px;
  }
  table {
    margin-top: 0px;
  }
}
</style>
