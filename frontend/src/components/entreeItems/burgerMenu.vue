<template>
  <div class="container text-center">
    <h1>Burger Menu</h1>
    <router-link class="btn btn-secondary mx-1 my-1" to="/sideMenu">All American Cheese Burger</router-link>
    <router-link class="btn btn-secondary mx-1 my-1" to="/sideMenu">Bacon Cheese Burger</router-link>
    <router-link class="btn btn-secondary mx-1 my-1" to="/sideMenu">Veggie Burger</router-link>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

@Component
export default class burgerMenu extends Vue {}
</script>

<style scoped>
h1 {
  text-align: center;
}
</style>
