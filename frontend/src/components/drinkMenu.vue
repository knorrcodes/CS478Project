<template>
  <div class="container text-center">
    <h1>Drink Menu</h1>
    <router-link class="btn btn-secondary mx-1 my-1" to="/mainMenu">Coke</router-link>
    <router-link class="btn btn-secondary mx-1 my-1" to="/mainMenu">Sprite</router-link>
    <router-link class="btn btn-secondary mx-1 my-1" to="/mainMenu">Dr. Pepper</router-link>
    <router-link class="btn btn-secondary mx-1 my-1" to="/mainMenu">Lemonade</router-link>
    <router-link class="btn btn-secondary mx-1 my-1" to="/mainMenu">Root Beer</router-link>
    <router-link class="btn btn-secondary mx-1 my-1" to="/mainMenu">Tea</router-link>
    <router-link class="btn btn-secondary mx-1 my-1" to="/mainMenu">Coffee</router-link>
    <router-link class="btn btn-secondary mx-1 my-1" to="/mainMenu">Water</router-link>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

@Component
export default class drinkMenu extends Vue {}
</script>

<style scoped>
h1 {
  text-align: center;
}
</style>