<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<style>
* {
  box-sizing: border-box;
}
.introduction-page {
  width: 80%;
  margin-left: auto;
  margin-right: auto;
}
table {
  border-collapse: collapse;
  height: max-content;
  margin-left: auto;
  margin-right: auto;
}
.board-td__white {
  background-color: white;
  width: 100%;
  height: 100%;
}
.board-td__blue {
  background-color: lightblue;
  width: 100%;
  height: 100%;
}
.board-td__darkblue {
  background-color: darkblue;
  width: 100%;
  height: 100%;
}
.board-td__gray {
  background-color: darkgray;
  width: 100%;
  height: 100%;
}
</style>
