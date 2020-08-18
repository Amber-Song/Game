<template>
  <div class="introduction-page">
    <h1>
      <router-link :to="{name: 'Home'}" class="link__none-style">Games</router-link>-- Notice
    </h1>

    <table>
      <tr>
        <td>August 17, 2020</td>
        <td>
          <strong>Monopoly</strong> local PC version is available.
          <span class="update">update</span>
        </td>
      </tr>

      <tr>
        <td>July 26, 2020</td>
        <td>
          <strong>Rotating puzzle</strong> local PC and mobile version is available.
        </td>
      </tr>

      <tr>
        <td>July 25, 2020</td>
        <td>
          <strong>Tic Tac Toe box version</strong> online PC and mobile version is available.
          <br>
          <strong>Tic Tac Toe box version</strong> local PC and mobile version is available.
        </td>
      </tr>

      <tr>
        <td>July 18, 2020</td>
        <td>
          <strong>Seek for airplane's head</strong> local PC and mobile version is available.
        </td>
      </tr>

      <tr>
        <td>July 14, 2020</td>
        <td>
          <strong>Seek for airplane's head</strong> online mobile version is available.
        </td>
      </tr>

      <tr>
        <td>June 16, 2020</td>
        <td>
          <strong>Seek for airplane's head</strong> online PC version is available.
        </td>
      </tr>
    </table>
  </div>
</template>

<style scoped>
table {
  margin-left: 0px;
  line-height: 160%;
}
tr {
  margin-top: 20px;
  vertical-align: top;
}
td:first-child {
  width: 130px;
}
.update {
  background-color: cornflowerblue;
  border: 1px solid cornflowerblue;
  border-radius: 4px;
  color: white;
}
</style>
