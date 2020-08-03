<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<style>
@import url("https://fonts.googleapis.com/css2?family=Neucha&display=swap");
@import url("https://fonts.googleapis.com/css2?family=Aladin&display=swap");
@import url("https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css");

* {
  box-sizing: border-box;
}
.introduction-page {
  width: 80%;
  max-width: 1350px;
  margin-left: auto;
  margin-right: auto;
  font-family: "Neucha", sans-serif;
}
.introduction-describe {
  font-family: "Neucha", sans-serif;
  font-size: 1.25em;
  margin: 30px;
  margin-top: 0px;
}
.introduction-button {
  display: block;
  font-family: "Neucha", sans-serif;
  font-size: 1.2em;
  margin: 40px auto;
  padding: 4px 10px 0 10px;
  border-radius: 3px;
}
.introduction-button:hover a {
  color: white;
}
button:hover {
  background-color: #003bac;
  border-top: 2px solid #608cdf;
  border-left: 2px solid #608cdf;
  border-bottom: 2px solid #002a7b;
  border-right: 2px solid #002a7b;
  color: white;
}

.game-page {
  width: 90%;
  min-width: 270px;
  max-width: 1350px;
  margin-left: auto;
  margin-right: auto;
  font-family: "Neucha", sans-serif;
}
.game-notice {
  background-color: #c6e2ff;
  background-image: linear-gradient(to right, #c6e2ff, white);
  margin-bottom: 40px;
}
.game-content {
  width: 100%;
  text-align: center;
  font-size: 1.25em;
  font-weight: bold;
  margin-bottom: 10px;
}
.game-congratulation {
  color: red;
}
.game-player {
  font-size: 1em;
  font-weight: normal;
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
h2 {
  font-family: "Aladin", cursive;
  font-weight: bold;
  margin: 0px;
  margin-bottom: 10px;
  margin-top: 10px;
}
table {
  border-collapse: collapse;
  height: max-content;
  margin-left: auto;
  margin-right: auto;
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
input {
  font-family: "Neucha", sans-serif;
  font-size: 1em;
  width: 50px;
  padding: 6px 0 0 5px;
}

/* Airplane */
.airplane-example__td {
  width: 30px;
  height: 30px;
  background-color: darkgray;
  border: 1px solid black;
  border-collapse: collapse;
}
.airplane__white {
  background-color: white;
  width: 100%;
  height: 100%;
}
.airplane__blue {
  background-color: lightblue;
  width: 100%;
  height: 100%;
}
.airplane__darkblue {
  background-color: darkblue;
  width: 100%;
  height: 100%;
}
.airplane__gray {
  background-color: darkgray;
  width: 100%;
  height: 100%;
}

/* Rotate puzzle rectangle */
.rotate-puzzle__example-td {
  width: 30px;
  height: 30px;
  border: 1px solid black;
  border-collapse: collapse;
}
.rotate-puzzle__red {
  background-color: #ff3e3e;
}
.rotate-puzzle__yellow {
  background-color: #fafd51;
}
.rotate-puzzle__green {
  background-color: #77ec77;
}
.rotate-puzzle__blue {
  background-color: #4f4fff;
}

/* Rotate puzzle triangle */
.triangle_example {
  display: inline-block;
  margin-top: 4px;
  margin-right: -10px;
  width: 0;
  height: 0;
  border-left: 15px solid transparent;
  border-right: 15px solid transparent;
  border-bottom: 15px solid white;
  vertical-align: top;
}
.triangle_rotate {
  transform: rotate(180deg);
}
.triangle_example_light-green {
  border-bottom: 30px solid #96ceb4;
}
.triangle_example_light-yellow {
  border-bottom: 30px solid #ffeead;
}
.triangle_example_green {
  border-bottom: 30px solid #3b9a9c;
}
.triangle_example_orange {
  border-bottom: 30px solid #ffad60;
}
.triangle_example_purple {
  border-bottom: 30px solid #a696c8;
}
.triangle_example_pink {
  border-bottom: 30px solid #fad3cf;
}

/* icon */
.title__icon {
  width: 60px;
  text-align: center;
}
.icon__airplane {
  transform: rotate(-45deg);
  color: #011962;
}
.icon__bullborn {
  margin: 0 15px;
  color: #011962;
}
.icon__scissor {
  transform: rotate(45deg);
  color: #c16400;
}

@media (max-width: 700px) {
  .introduction-page {
    width: 100%;
  }
  .introduction-describe {
    font-size: 1em;
  }
  .introduction-button {
    font-size: 1em;
    margin-top: 10px;
    margin-bottom: 20px;
  }

  .game-page {
    width: 100%;
  }
  .game-notice {
    margin-bottom: 10px;
  }
  .game-content {
    font-size: 1em;
  }

  h1 {
    font-family: "Aladin", cursive;
    font-size: 1.5em;
  }
}
</style>
