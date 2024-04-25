<template>
  <div class="container">
    <div class="box">
      <div class="front"></div>
      <div class="back"></div>
      <div class="left"></div>
      <div class="right"></div>
      <div class="top"></div>
      <div class="bottom"></div>
    </div>
  </div>
</template>

<style scoped>
.container {
  height: 80px;
  perspective: 200px;
}
.box {
  height: 80px;
  width: 80px;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: rotateX(-10deg) rotateY(10deg);
  transform-style: preserve-3d;
}
.box div {
  height: 80px;
  width: 80px;
  text-align: center;
  position: absolute;
}
.front {
  background: rgba(103, 93, 23, 1);
  transform: translateZ(40px);
}
.back {
  background: rgba(103, 93, 23, 1);
  transform: translateZ(-40px);
}
.left {
  background: rgba(103, 93, 23, 1);
  transform: translateX(-40px) rotateY(-90deg);
}
.right {
  background: rgba(103, 93, 23, 1);
  transform: translateX(40px) rotateY(90deg);
}
.top {
  background: rgba(103, 93, 23, 1);
  transform: translateY(-40px) rotateX(90deg);
}
.bottom {
  background: rgba(103, 93, 23, 0);
  transform: translateY(40px) rotateX(-90deg);
}
</style>
