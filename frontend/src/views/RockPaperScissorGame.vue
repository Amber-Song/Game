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
          <div v-for="(card, index) in player1Cards" :key="index">
            <div v-if="thisPlayer == 'player1'" v-on:click="chooseCard(index)">
              <div v-if=" card === 'rock' && index != player1Choose" class="fa fa-hand-rock-o card"></div>
              <div
                v-else-if=" card === 'paper' && index != player1Choose"
                class="fa fa-hand-paper-o card"
              ></div>
              <div
                v-else-if=" card === 'scissor' && index != player1Choose"
                class="fa fa-hand-scissors-o card"
              ></div>
            </div>
            <div v-else>
              <div class="card"></div>
            </div>
          </div>
        </div>
      </div>

      <div class="title-center">
        <div>
          <strong>Round {{ round }}</strong>
        </div>
        <div
          v-if="thisPlayer == 'player1' && player1Choose == -1"
        >Please choose one card and click decide button to place on the table!</div>
        <div
          v-if="thisPlayer == 'player2' && player2Choose == -1"
        >Please choose one card and click decide button to place on the table!</div>
        <div v-if="announceRoundWinner != ''">{{ announceRoundWinner }}</div>

        <div v-if="timeCount >= 0 && timeCount < 10" class="time">
          Timer:
          <strong>{{ timeCount }}</strong>
        </div>
        <div
          v-if="winner != ''"
          class="announce__winner"
        >{{ winner }} win this game! Congratulation!</div>
        <div class="competition__table">
          <div class="card__empty">
            <div v-if="player1Choose != -1" class="card__placed">
              <div v-if="player1Flip">
                <div v-if=" player1Card === 'rock'" class="fa fa-hand-rock-o card__competition"></div>
                <div v-if=" player1Card === 'paper'" class="fa fa-hand-paper-o card__competition"></div>
                <div
                  v-if=" player1Card === 'scissor'"
                  class="fa fa-hand-scissors-o card__competition"
                ></div>
              </div>
            </div>
          </div>

          <div class="card__empty">
            <div v-if="player2Choose != -1" class="card__placed">
              <div v-if="player2Flip">
                <div v-if=" player2Card === 'rock'" class="fa fa-hand-rock-o card__competition"></div>
                <div v-if=" player2Card === 'paper'" class="fa fa-hand-paper-o card__competition"></div>
                <div
                  v-if=" player2Card === 'scissor'"
                  class="fa fa-hand-scissors-o card__competition"
                ></div>
              </div>
            </div>
          </div>
        </div>
        <button v-on:click="clearTime()">
          Back to introduction page to
          <strong>restart</strong>.
        </button>
      </div>

      <div>
        <div class="title-center">Player2</div>
        <div class="collection">
          <div v-for="(card, index) in player2Cards" :key="index">
            <div v-if="thisPlayer == 'player2'" v-on:click="chooseCard(index)">
              <div v-if=" card === 'rock' && index != player2Choose" class="fa fa-hand-rock-o card"></div>
              <div
                v-else-if=" card === 'paper' && index != player2Choose"
                class="fa fa-hand-paper-o card"
              ></div>
              <div
                v-else-if=" card === 'scissor' && index != player2Choose"
                class="fa fa-hand-scissors-o card"
              ></div>
            </div>
            <div v-else>
              <div class="card"></div>
            </div>
          </div>
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
      winner: "",
      player1Cards: [],
      player2Cards: [],
      thisPlayer: "",
      player1: "",
      player2: "",
      timeStop: null,
      timeCount: 10,
      player1Choose: -1,
      player2Choose: -1,
      player1Card: "",
      player2Card: "",
      player1Flip: true,
      player2Flip: true,
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
        .get(`${this.$hostname}/Game/api/RockPaperScissor/Game/wait`, {
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
            window.setTimeout(this.roundStart, 3000);
          }
        })
        .catch(err => {
          this.errors.push(err);
        });
    },

    clearTime() {
      clearTimeout(this.timeStop);
    },

    roundStart() {
      if (this.winner == "") {
        window.setTimeout(this.roundStart, 17000);
      } else {
        return;
      }

      this.timeCount = 10;

      // Wait for 10 seconds for the answer
      this.TimerReduceOne();

      // After waiting, post the result, flip card
      window.setTimeout(this.postChosenCard, 10000);
      // Wait for a second to get result, show both cards and result
      window.setTimeout(this.getCardResult, 12000);
      // Get round result and update collection
      window.setTimeout(this.getRoundResult, 16000);
      // Check if game end. If no, start the next round. If yes, announce winner
    },

    TimerReduceOne() {
      this.timeCount = this.timeCount - 1;
      if (this.timeCount < 0) {
        clearTimeout(this.timeStop);
        return;
      }
      this.timeStop = window.setTimeout(this.TimerReduceOne, 1000);
    },

    chooseCard(index) {
      if (this.thisPlayer == "player1") {
        this.player1Choose = index;
        this.player1Card = this.player1Cards[index];
      }
      if (this.thisPlayer == "player2") {
        this.player2Choose = index;
        this.player2Card = this.player2Cards[index];
      }
    },

    // Post the chosen card to server
    postChosenCard() {
      if (this.thisPlayer == "player1") {
        this.axios
          .post(
            `${this.$hostname}/Game/api/RockPaperScissor/Game/room`,
            {
              Card1: this.player1Card,
              Card1Index: this.player1Choose
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
        this.axios
          .post(
            `${this.$hostname}/Game/api/RockPaperScissor/Game/room`,
            {
              Card2: this.player2Card,
              Card2Index: this.player2Choose
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
      this.flipCard();
    },

    // Get the chosen card from the server
    getCardResult() {
      this.axios
        .get(`${this.$hostname}/Game/api/RockPaperScissor/Game/room`, {
          withCredentials: true,
          params: {
            room: this.getRoom
          }
        })
        .then(response => {
          // Get data
          this.player1Card = response.data.Card1;
          this.player2Card = response.data.Card2;
          this.player1Choose = response.data.Card1Index;
          this.player2Choose = response.data.Card2Index;
          this.flipCard();
          this.checkRoundWin();
        });
    },

    // Get the collection change and reset the chosen card
    getRoundResult() {
      this.axios
        .get(`${this.$hostname}/Game/api/RockPaperScissor/Game/roundend`, {
          withCredentials: true,
          params: {
            room: this.getRoom
          }
        })
        .then(response => {
          this.round = response.data.Round;
          this.player1Cards = response.data.Collection1;
          this.player2Cards = response.data.Collection2;
          this.player1Choose = -1;
          this.player2Choose = -1;
          this.player1Card = "";
          this.player2Card = "";
          this.announceRoundWinner = "";
          this.checkWinner();
        });
    },

    flipCard() {
      this.player1Flip = !this.player1Flip;
      this.player2Flip = !this.player2Flip;
    },

    checkRoundWin() {
      if (this.player1Card == this.player2Card) {
        this.announceRoundWinner = "Two players are draw!";
        return;
      }

      switch (this.player1Card) {
        case "rock":
          switch (this.player2Card) {
            case "paper":
              // player2 win
              this.announceRoundWinner = "Player2 win this round!";
              return;
            case "scissor":
              // player1 win
              this.announceRoundWinner = "Player1 win this round!";
              return;
            default:
              this.announceRoundWinner =
                "Player2 didn't choose cards. Player1 win this round!";
              return;
          }
        case "paper":
          switch (this.player2Card) {
            case "rock":
              // player1 win
              this.announceRoundWinner = "Player1 win this round!";
              return;
            case "scissor":
              // player2 win
              this.announceRoundWinner = "Player2 win this round!";
              return;
            default:
              this.announceRoundWinner =
                "Player2 didn't choose cards. Player1 win this round!";
              return;
          }
        case "scissor":
          switch (this.player2Card) {
            case "rock":
              // player2 win
              this.announceRoundWinner = "Player2 win this round!";
              return;
            case "paper":
              // player1 win
              this.announceRoundWinner = "Player1 win this round!";
              return;
            default:
              this.announceRoundWinner =
                "Player2 didn't choose cards. Player1 win this round!";
              return;
          }
        default:
          this.announceRoundWinner =
            "Player1 didn't choose cards. Player2 win this round!";
          return;
      }
    },

    checkWinner() {
      if (this.player1Cards.length == 0) {
        this.winner = "Player1";
      }
      if (this.player2Cards.length == 0) {
        this.winner = "Player2";
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
.announce__winner {
  color: red;
  font-size: 1.5em;
  margin-top: 5em;
}
.time {
  margin-top: 1em;
  font-size: 1.5em;
}

button {
  font-family: "Neucha", sans-serif;
  font-size: 1em;
  margin: 15px 0 0 15px;
  padding: 4px 20px 0 20px;
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
