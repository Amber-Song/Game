<template>
  <div class="introduction-page">
    <h1>Games</h1>
    <div class="page-notice">
      * To play the game, one player click play now button which will jump to the game page
      with a special room number. And the player who enters the room first, copy and send the URL
      to the other player. While the other player enters the URL, the player will automatically enter the room.
    </div>
    <ul>
      <li>
        <router-link
          :to="{name: 'FindAirplaneIntroduction'}"
          class="link__none-style"
        >Find Airplane's Head</router-link>
      </li>
    </ul>
  </div>
</template>

<style scoped>
.page-notice {
  border: 2px dotted transparent;
  border-image: linear-gradient(to top, #f9a507, #00b176);
  border-image-slice: 10;
  padding: 10px;
}
</style>
