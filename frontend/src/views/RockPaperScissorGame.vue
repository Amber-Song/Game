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
            <div v-if="thisPlayer == 'player1'">
              <div
                v-if=" card === 'rock'"
                class="fa fa-hand-rock-o card"
                :class="{ 'card-chose': index === player1Choose, 'card__hover': index === player1Choose}"
              ></div>
              <div
                v-else-if=" card === 'paper'"
                class="fa fa-hand-paper-o card"
                :class="{ 'card-chose': index === player1Choose, 'card__hover': index === player1Choose}"
              ></div>
              <div
                v-else
                class="fa fa-hand-scissors-o card"
                :class="{ 'card-chose': index === player1Choose, 'card__hover': index === player1Choose}"
              ></div>
            </div>
            <div v-else>
              <div class="card"></div>
            </div>
          </div>
        </div>
        <button v-on:click="decide()" v-if="thisPlayer === 'player1'">Decide</button>
      </div>

      <div class="title-center">
        <strong>Round {{ round }}</strong>
        <br>please choose one card and place on the table!
        <br>
        Timer: {{ timeCount }}
        <div v-if="announceRoundWinner != ''">{{ announceRoundWinner }}</div>
        <div v-if="win != ''">{{ win }} win this game! Congratulation!</div>
        <div class="competition__table">
          <div class="card__empty" :class="{ 'card__placed': player1ChooseDecide != -1}">
            <div
              v-if=" player1Flip && player1Card === 'rock'"
              class="fa fa-hand-rock-o card__competition"
            ></div>
            <div
              v-if=" player1Flip && player1Card === 'paper'"
              class="fa fa-hand-paper-o card__competition"
            ></div>
            <div
              v-if=" player1Flip && player1Card === 'scissor'"
              class="fa fa-hand-scissors-o card__competition"
            ></div>
          </div>
          <div class="card__empty" :class="{ 'card__placed': player2ChooseDecide != -1}">
            <div
              v-if=" player2Flip && player2Card === 'rock'"
              class="fa fa-hand-rock-o card__competition"
            ></div>
            <div
              v-if=" player2Flip && player2Card === 'paper'"
              class="fa fa-hand-paper-o card__competition"
            ></div>
            <div
              v-if=" player2Flip && player2Card === 'scissor'"
              class="fa fa-hand-scissors-o card__competition"
            ></div>
          </div>
        </div>
      </div>

      <div>
        <div class="title-center">Player2</div>
        <div class="collection">
          <div v-for="(card, index) in player2Cards" :key="index" v-on:click="chooseCard(index)">
            <div v-if="thisPlayer == 'player2'">
              <div
                v-if=" card === 'rock'"
                class="fa fa-hand-rock-o card"
                :class="{ 'card-chose': index === player2Choose}"
              ></div>
              <div
                v-else-if=" card === 'paper'"
                class="fa fa-hand-paper-o card"
                :class="{ 'card-chose': index === player2Choose}"
              ></div>
              <div
                v-else
                class="fa fa-hand-scissors-o card"
                :class="{ 'card-chose': index === player2Choose}"
              ></div>
            </div>
            <div v-else>
              <div class="card"></div>
            </div>
          </div>
        </div>
        <button v-on:click="decide()" v-if="thisPlayer === 'player2'">Decide</button>
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
      timeCount: 10,
      player1Choose: -1,
      player2Choose: -1,
      player1ChooseDecide: -1,
      player2ChooseDecide: -1,
      player1Card: "",
      player2Card: "",
      player1Flip: false,
      player2Flip: false,
      announceRoundWinner: ""
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
            this.timeStop = window.setTimeout(this.queryServer, 1000);
          }
        })
        .catch(err => {
          this.errors.push(err);
        });
    },

    chooseCard(index) {
      if (this.thisPlayer == "player1") {
        this.player1Choose = index;
      }
      if (this.thisPlayer == "player2") {
        this.player2Choose = index;
      }
    },

    decide() {
      if (this.thisPlayer == "player1") {
        this.player1ChooseDecide = this.player1Choose;
        this.axios
          .post(
            `${this.$hostname}/api/RockPaperScissor/Game/room`,
            {
              Card1: this.player1Cards[this.player1ChooseDecide],
              Card1Index: this.player1ChooseDecide
            },
            {
              withCredentials: true,
              params: {
                room: this.getRoom
              }
            }
          )
          .catch(err => {
            this.errors.push(err);
          });
      }
      if (this.thisPlayer == "player2") {
        this.player2ChooseDecide = this.player2Choose;
        this.axios
          .post(
            `${this.$hostname}/api/RockPaperScissor/Game/room`,
            {
              Card2: this.player2Cards[this.player2ChooseDecide],
              Card2Index: this.player2ChooseDecide
            },
            {
              withCredentials: true,
              params: {
                room: this.getRoom
              }
            }
          )
          .catch(err => {
            this.errors.push(err);
          });
      }
    },

    queryServer() {
      this.axios
        .get(`${this.$hostname}/api/RockPaperScissor/Game/room`, {
          withCredentials: true,
          params: {
            room: this.getRoom
          }
        })
        .then(response => {
          // Get data
          this.player1Card = response.data.Card1;
          this.player2Card = response.data.Card2;
          this.player1ChooseDecide = response.data.Card1Index;
          this.player2ChooseDecide = response.data.Card2Index;
        });

      this.timeStop = window.setTimeout(this.queryServer, 1000);
      this.timeCount = this.timeCount - 1;
      if (this.timeCount <= 0) {
        clearTimeout(this.timeStop);
        window.setTimeout(this.flipCard, 1000);
        window.setTimeout(this.checkRoundWin, 5000);
        window.setTimeout(this.instructEnd, 5000);
      }
    },

    instructEnd() {
      // Call the server to renew the collections
      this.axios
        .get(`${this.$hostname}/api/RockPaperScissor/Game/roundend`, {
          withCredentials: true,
          params: {
            room: this.getRoom
          }
        })
        .then(response => {
          console.log(response);
          this.round = response.data.Round;
          this.player1Cards = response.data.Collection1;
          this.player2Cards = response.data.Collection2;
          this.timeCount = 10;
          this.player1Choose = -1;
          this.player2Choose = -1;
          this.player1ChooseDecide = -1;
          this.player2ChooseDecide = -1;
          this.player1Card = "";
          this.player2Card = "";
          this.player1Flip = false;
          this.player2Flip = false;
          this.announceRoundWinner = "";
        });

      this.checkRoundWin();
      if (this.win != "") {
        clearTimeout(this.timeStop);
      } else {
        clearTimeout(this.timeStop);
        this.timeStop = window.setTimeout(this.queryServer, 1000);
      }
    },

    flipCard() {
      this.player1Flip = true;
      this.player2Flip = true;
    },

    checkRoundWin() {
      switch (this.player1Card) {
        case "rock":
          switch (this.player2Card) {
            case "rock":
              // draw
              this.announceRoundWinner = "Two players are draw!";
              break;
            case "paper":
              // player2 win
              this.announceRoundWinner = "Player2 win this round!";
              break;
            case "scissor":
              // player1 win
              this.announceRoundWinner = "Player1 win this round!";
              break;
            default:
              this.announceRoundWinner =
                "Player2 didn't choose cards. Player1 win this round!";
              break;
          }
          break;
        case "paper":
          switch (this.player2Card) {
            case "rock":
              // player1 win
              this.announceRoundWinner = "Player1 win this round!";
              break;
            case "paper":
              // draw
              this.announceRoundWinner = "Two players are draw!";
              break;
            case "scissor":
              // player2 win
              this.announceRoundWinner = "Player2 win this round!";
              break;
            default:
              this.announceRoundWinner =
                "Player2 didn't choose cards. Player1 win this round!";
              break;
          }
          break;
        case "scissor":
          switch (this.player2Card) {
            case "rock":
              // player2 win
              this.announceRoundWinner = "Player2 win this round!";
              break;
            case "paper":
              // player1 win
              this.announceRoundWinner = "Player1 win this round!";
              break;
            case "scissor":
              // draw
              this.announceRoundWinner = "Two players are draw!";
              break;
            default:
              this.announceRoundWinner =
                "Player2 didn't choose cards. Player1 win this round!";
              break;
          }
          break;
        default:
          if (this.player2Card == "") {
            this.announceRoundWinner = "Two players are draw!";
          } else {
            this.announceRoundWinner =
              "Player1 didn't choose cards. Player2 win this round!";
          }
          break;
      }
    },

    checkWinner() {
      if (this.player1Cards.length == 0) {
        this.win = "Player1";
      }
      if (this.player2Cards.length == 0) {
        this.win = "Player2";
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
.card__hover:hover {
  box-shadow: 2px 2px 1px #fbd795;
}
.card-chose {
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
  border: 5px solid #ffdb98;
  border-radius: 3px;
}
.card__empty {
  width: 200px;
  height: 300px;
}
.card__competition {
  font-size: 130px;
  padding-top: 80px;
  color: #d57005;
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
