<template>
  <div class="background" id="app">
    <img class="logoImage" src="../../docs/biglogo.png" alt="big team logo" />
    <div class="row">
      <div class="col-4 px-4">
        <table-order />
      </div>
      <div class="col-8">
        <router-view />
      </div>
    </div>
  </div>
</template>

<script>
import loginScreen from "./components/loginScreen.vue";
import mainMenu from "./components/mainMenu.vue";
import tableOrder from "./components/tableOrder.vue";
import ApiClass from "./api";

const axios = require("axios");

export default {
  name: "app",
  components: {
    loginScreen,
    mainMenu,
    tableOrder
  },
  data() {
    return {
      apiData: null
    };
    currentPath = window.location.path;
    showTicket = true;
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
  height: 100%;
  min-height: 100vh;
}

.logoImage {
  height: 200;
  display: block;
  margin-left: auto;
  margin-right: auto;
}
</style>
