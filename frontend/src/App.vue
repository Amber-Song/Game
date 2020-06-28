<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<style>
@import url("https://fonts.googleapis.com/css2?family=Neucha&display=swap");
@import url("https://fonts.googleapis.com/css2?family=Aladin&display=swap");

* {
  box-sizing: border-box;
}
.page {
  width: 90%;
  max-width: 1350px;
  margin-left: auto;
  margin-right: auto;
  font-family: "Neucha", sans-serif;
}
.introduction-page {
  width: 80%;
  max-width: 1350px;
  margin-left: auto;
  margin-right: auto;
  font-family: "Neucha", sans-serif;
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
.link__none-style,
.link__none-style:visited {
  text-decoration: none;
  color: black;
}
h1 {
  font-family: "Aladin", cursive;
  font-size: 2.5em;
}
.icon__airplane {
  rotate: -45deg;
  color: #011962;
}
.icon__bullborn {
  margin: 0 15px;
  color: #011962;
}
</style>
