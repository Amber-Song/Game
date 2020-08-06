<template>
  <div id="airplaneA">
    <table>
      <tr v-for="(array, indexarray) in airplaneA" :key="indexarray">
        <td v-for="(block, indexblock) in array" :key="indexblock" class="board-example">
          <div
            v-bind:class="{'airplane__white': block === 3, 'airplane__blue': block === 4, 'airplane__darkblue':block === 5 }"
          ></div>
        </td>
      </tr>
    </table>
  </div>
</template>

<script>
export default {
  computed: {
    airplaneA() {
      return this.$store.state.airplaneA;
    }
  }
};
</script>

<style scoped>
.board-example {
  width: 15px;
  height: 15px;
  background-color: darkgray;
  border: 1px solid black;
  border-collapse: collapse;
}
</style>
