<template>
  <div>
    <div class="container text-center">
      <h1>Main Menu</h1>
      <div class="row">
        <div class="col" align="center">
          <button type="button" class="btn btn-secondary mx-1" @click="clickedChicken">Chicken</button>
          <button type="button" class="btn btn-secondary mx-1" @click="clickedSteak">Steak</button>
          <button type="button" class="btn btn-secondary mx-1" @click="clickedFish">Fish</button>
          <button type="button" class="btn btn-secondary mx-1" @click="clickedBurgers">Burgers</button>
          <button type="button" class="btn btn-secondary mx-1" @click="clickedApatizers">Apatizers</button>
          <button type="button" class="btn btn-secondary mx-1" @click="clickedDesserts">Desserts</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";

@Component
export default class mainMenu extends Vue {
  private clickedChicken() {
    console.log("chicken clicked");
    this.$router.push({ name: "sideMenu" });
  }
  private clickedSteak() {
    console.log("Clicked Steak");
    this.$router.push({ name: "steakMenu" });
  }
  private clickedFish() {
    console.log("Clicked Fish");
    this.$router.push({ name: "fishMenu" });
  }
  private clickedBurgers() {
    console.log("Clicked Burgers");
    this.$router.push({ name: "burgerMenu" });
  }
  private clickedApatizers() {
    console.log("Clicked Apatizers");
    this.$router.push({ name: "apatizerMenu" });
  }
  private clickedDesserts() {
    console.log("Clicked Deserts");
    this.$router.push({ name: "dessertMenu" });
  }
}
</script>

<style scoped>
</style>