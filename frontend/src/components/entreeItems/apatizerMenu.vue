<template>
  <div class="container text-center">
    <h1>Apatizer Menu</h1>
    <router-link class="btn btn-secondary mx-1 my-1" to="/drinkMenu">Fried Pickles</router-link>
    <router-link class="btn btn-secondary mx-1 my-1" to="/drinkMenu">Mozz Sticks</router-link>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

@Component
export default class apatizerMenu extends Vue {}
</script>

<style scoped>
h1 {
  text-align: center;
}
</style>
