<template>
  <div class="sideMenu component">
    <h1>Sides Menu</h1>
    <div class="row justify-content-center">
      <router-link class="btn btn-secondary mx-1" to="/drinkMenu">French Fries</router-link>
      <router-link class="btn btn-secondary mx-1" to="/drinkMenu">Baked Potato</router-link>
      <router-link class="btn btn-secondary mx-1" to="/drinkMenu">Broccoli</router-link>
      <router-link class="btn btn-secondary mx-1" to="/drinkMenu">Asparagus</router-link>
      <router-link class="btn btn-secondary mx-1" to="/drinkMenu">Rice</router-link>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

@Component
export default class sideMenu extends Vue {}
</script>

<style scoped>
h1 {
  text-align: center;
}
</style>
