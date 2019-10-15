<template>
  <div>
    <h1>Drink Menu</h1>
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