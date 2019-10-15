<template>
  <div class="container text-center">
    <h1>Dessert Menu</h1>
    <router-link class="btn btn-secondary mx-1" to="/drinkMenu">Epic Brownie</router-link>
    <router-link class="btn btn-secondary mx-1" to="/drinkMenu">Fried Ice Cream</router-link>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

@Component
export default class dessertMenu extends Vue {}
</script>

<style scoped>
h1 {
  text-align: center;
}
</style>
