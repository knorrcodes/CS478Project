<template>
  <div class="background" id="app">
    <div>
      <img class="logoImage" src="../../docs/biglogo.png" alt="biglogo" />
    </div>
    <div id="login"></div>
    <login-screen />
  </div>
</template>

<script>
import loginScreen from "./components/loginScreen.vue";

export default {
  name: "app",
  components: {
    loginScreen
  }
};
</script>

<style>
body {
  background-color: firebrick;
  height: 100%;
}

.logoImage {
  height: 200;
  display: block;
  margin-left: auto;
  margin-right: auto;
}
</style>
