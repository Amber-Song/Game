<template>
  <div id="gameNotice" class="game-notice">
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
</template>

<script>
export default {
  props: {
    thisPlayer: String,
    player2: String
  }
};
</script>
