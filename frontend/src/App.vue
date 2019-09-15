<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">Home</router-link>|
      <router-link to="/about">About</router-link>
      <div>Hello there</div>
      <comp1 />
    </div>
  </div>
</template>

<script>
import comp1 from "./components/comp1.vue";

export default {
  name: "app",
  components: {
    comp1
  }
};
</script>

<style>
</style>
