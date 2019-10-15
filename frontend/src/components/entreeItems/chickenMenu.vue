<template>
  <div class="containter text-center">
    <h1>Chicken Menu</h1>
    <router-link class="btn btn-secondary mx-1" to="/sideMenu">Chicken Tender Dinner</router-link>
    <router-link class="btn btn-secondary mx-1" to="/sideMenu">Tariyaki Chicken</router-link>
    <router-link class="btn btn-secondary mx-1" to="/sideMenu">Grilled Chicken</router-link>
    <router-link class="btn btn-secondary mx-1" to="/sideMenu">Chicken Salad</router-link>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

@Component
export default class chickenMenu extends Vue {}
</script>

<style scoped>
h1 {
  text-align: center;
}
</style>
