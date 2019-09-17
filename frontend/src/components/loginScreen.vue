<template>
  <div>
    <div class="id1 container">
      <div class="row justify-content-center">
        <form action>
          <p class="text-center">Server Login</p>
          <input class="input1" id="serverCodeNumber" type="number" value />
          <br />
          <button class="input1" id="serverCodeButton" @click="serverCode()">Enter</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "loginScreen"
};

function serverCode() {}
</script>

<style scoped>
.id1 {
  text-align: center;
  justify-content: center;
}
.input1 {
  display: block;
  margin-left: auto;
  margin-right: auto;
}
</style>