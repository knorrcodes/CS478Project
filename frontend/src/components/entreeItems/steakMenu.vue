<template>
  <div class="containter text-center">
    <h1>Steak Menu</h1>
    <router-link class="btn btn-secondary mx-1" to="/sideMenu">6oz Filet</router-link>
    <router-link class="btn btn-secondary mx-1" to="/sideMenu">9oz Filet</router-link>
    <router-link class="btn btn-secondary mx-1" to="/sideMenu">Steak Salad</router-link>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

@Component
export default class steakMenu extends Vue {}
</script>

<style scoped>
h1 {
  text-align: center;
}

.centerContents {
  align-content: center;
}
</style>
