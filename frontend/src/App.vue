<template>
  <div class="background" id="app">
    <img class="logoImage" src="../../docs/biglogo.png" alt />
    <router-view />
  </div>
</template>

<script>
import loginScreen from "./components/loginScreen.vue";
import mainMenu from "./components/mainMenu.vue";
import ApiClass from "./api";

const axios = require("axios");

export default {
  name: "app",
  components: {
    loginScreen,
    mainMenu
  },
  data() {
    return {
      apiData: null
    };
  }
  /*
  mounted() {
    axios
      .get("https://api.coindesk.com/v1/bpi/currentprice.json")
      .then(results => (this.apiData = results))
      .catch(err => console.log(this.err));
  } */
};
</script>

<style>
.background {
  background-color: rgb(255, 111, 111);
  height: 100vh;
}

.logoImage {
  height: 200;
  display: block;
  margin-left: auto;
  margin-right: auto;
}
</style>
