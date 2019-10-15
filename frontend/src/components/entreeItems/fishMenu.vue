<template>
  <div class="container text-center">
    <h1>Fish Menu</h1>
    <router-link class="btn btn-secondary mx-1 my-1" to="/sideMenu">Wild Salmon</router-link>
    <router-link class="btn btn-secondary mx-1 my-1" to="/sideMenu">Catfish Dinner</router-link>
    <router-link class="btn btn-secondary mx-1 my-1" to="/sideMenu">Grilled Shrimp</router-link>
    <router-link class="btn btn-secondary mx-1 my-1" to="/sideMenu">Lobster Tail</router-link>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

@Component
export default class fishMenu extends Vue {}
</script>

<style scoped>
h1 {
  text-align: center;
}
</style>
