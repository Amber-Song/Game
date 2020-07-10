<template>
  <div class="page">
    <h1>
      <span class="fa fa-hand-scissors-o icon__scissor"></span>
      Rock Paper Scissors
    </h1>

    <div class="notice">
      <font-awesome-icon :icon="['fas', 'bullhorn']" class="icon__bullborn"/>
      <span v-if="thisPlayer == 'player1'">
        Welcome! You are
        <strong>player1</strong> !
      </span>
      <span v-else>Player1 entered the room.</span>
      <span v-if="player2 != ''">
        <span v-if="thisPlayer == 'player2'">
          Welcome! You are
          <strong>player2</strong> !
          <strong>Game start !</strong>
        </span>
        <span v-else>
          Player2 entered the room.
          <strong>Game start !</strong>
        </span>
      </span>
      <span v-else>Please wait for player2 !</span>
    </div>

    <!-- Here is the game -->
    <div class="competition">
      <div>
        <div class="title-center">Player1</div>
        <div class="collection">
          <div v-for="(card, index) in player1Cards" :key="index" v-on:click="chooseCard(index)">
            <div v-if=" card === 'rock'" class="fa fa-hand-rock-o card"></div>
            <div v-else-if=" card === 'paper'" class="fa fa-hand-paper-o card"></div>
            <div v-else class="fa fa-hand-scissors-o card"></div>
          </div>
          <button v-on:click="decide()">Decide</button>
        </div>
      </div>
      <div class="title-center">
        <strong>Round {{ round }}</strong>
        <br>please choose one card and place on the table!
        <br>Timer:
        <div class="competition__table">
          <div class="card__placed" v-if="player1ChooseDecide != -1"></div>
          <div class="card__placed" v-if="player2ChooseDecide != -1"></div>
        </div>
      </div>
      <div>
        <div class="title-center">Player2</div>
        <div class="collection">
          <div v-for="(card, index) in player2Cards" :key="index">
            <div v-if=" card === 'rock'" class="fa fa-hand-rock-o card"></div>
            <div v-else-if=" card === 'paper'" class="fa fa-hand-paper-o card"></div>
            <div v-else class="fa fa-hand-scissors-o card"></div>
          </div>
          <button>Decide</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    roomid: {
      type: String
    }
  },
  data: function() {
    return {
      round: 0,
      win: "",
      player1Cards: [],
      player2Cards: [],
      thisPlayer: "",
      player1: "",
      player2: "",
      timeStop: null,
      player1Choose: -1,
      player2Choose: -1,
      player1ChooseDecide: -1,
      player2ChooseDecide: -1
    };
  },
  computed: {
    getRoom() {
      return this.$store.state.roomid;
    }
  },
  mounted: function() {
    this.$store.commit("getRoom", { roomid: this.roomid });
    this.loadData();
  },
  methods: {
    loadData() {
      this.timeStop = window.setTimeout(this.loadData, 1000);

      this.axios
        .get(`${this.$hostname}/api/RockPaperScissor/Game/room`, {
          withCredentials: true,
          params: { room: this.getRoom }
        })
        .then(response => {
          this.player1Cards = response.data.Collection1;
          this.player2Cards = response.data.Collection2;
          this.thisPlayer = response.data.ThisPlayer;
          this.player1 = response.data.Player1;
          this.player2 = response.data.Player2;
          this.round = response.data.Round;
          if (this.player1 != "" && this.player2 != "") {
            clearTimeout(this.timeStop);
          }
        })
        .catch(err => {
          this.errors.push(err);
        });
    },

    chooseCard(index) {
      if (this.thisPlayer == "player1") {
        this.player1Choose = index;
      } else {
        this.player2Choose = index;
      }
    },

    decide() {
      if (this.thisPlayer == "player1") {
        this.player1ChooseDecide = this.player1Choose;
      } else {
        this.player2ChooseDecide = this.player2Choose;
      }
    }
  }
};
</script>


<style scoped>
.notice {
  background-color: #c6e2ff;
  background-image: linear-gradient(to right, #c6e2ff, white);
  margin-bottom: 40px;
}
.title-center {
  font-family: "Aladin", cursive;
  text-align: center;
}

.competition {
  display: grid;
  grid-template-columns: 110px auto 110px;
}
.collection {
  display: flex;
  flex-flow: row wrap;
  justify-content: space-around;
}
.card {
  font-size: 2em;
  text-align: center;
  color: #d57005;
  border: 2px solid #ffdb98;
  border-radius: 3px;
  width: 46px;
  height: 64px;
  margin: 5px 0px;
  padding-top: 14px;
}
.card:hover {
  box-shadow: 2px 2px 1px #fbd795;
}
.competition__table {
  display: grid;
  grid-template-columns: 200px 200px;
  grid-column-gap: 40px;
  width: max-content;
  margin-left: auto;
  margin-right: auto;
}
.card__placed {
  width: 200px;
  height: 300px;
  border: 2px solid #ffdb98;
  border-radius: 3px;
}

button {
  font-family: "Neucha", sans-serif;
  font-size: 1em;
  margin: 15px 0;
  padding: 4px 10px 0 10px;
  border-radius: 3px;
  background-color: #003bac;
  border-top: 2px solid #608cdf;
  border-left: 2px solid #608cdf;
  border-bottom: 2px solid #002a7b;
  border-right: 2px solid #002a7b;
  color: white;
}
button:hover {
  background-color: #fdf0ab;
  border-top: 2px solid #fff5c3;
  border-left: 2px solid #fff5c3;
  border-bottom: 2px solid #eada86;
  border-right: 2px solid #eada86;
  color: black;
}
</style>
