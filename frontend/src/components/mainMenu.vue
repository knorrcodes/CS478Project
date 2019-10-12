<template>
  <div>
    <div class="container text-center">
      <h1>Main Menu</h1>
      <div class="row">
        <div class="col" align="center">
          <router-link class="btn btn-secondary mx-1" to="/chickenMenu">Chicken</router-link>
          <router-link class="btn btn-secondary mx-1" to="/sideMenu">Sides</router-link>
          <router-link class="btn btn-secondary mx-1" to="/steakMenu">Steak</router-link>
          <router-link class="btn btn-secondary mx-1" to="/fishMenu">Fish</router-link>
          <router-link class="btn btn-secondary mx-1" to="/burgerMenu">Burgers</router-link>
          <router-link class="btn btn-secondary mx-1" to="/apatizerMenu">Apatizers</router-link>
          <router-link class="btn btn-secondary mx-1" to="/dessertMenu">Desserts</router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";

@Component
export default class mainMenu extends Vue {}
</script>

<style scoped>
</style>
