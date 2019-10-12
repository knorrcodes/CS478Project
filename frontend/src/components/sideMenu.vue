<template>
  <div class="sideMenu component">
    <h1>Sides Menu</h1>
    <div class="row justify-content-center">
      <form action>
        <p class="text-center textFormat">Sides</p>
        <button
          type="button"
          class="btn btn-secondary mx-1"
          id="btnFrenchFries"
          @click="clickedFrenchFries"
        >French Fries</button>

        <button
          type="button"
          class="btn btn-secondary mx-1"
          id="btnPotato"
          @click="clickedBakedPotato"
        >Baked Potato</button>

        <button
          type="button"
          class="btn btn-secondary mx-1"
          id="btnBroccoli"
          @click="clickedBroccoli"
        >Broccoli</button>

        <button type="button" class="btn btn-secondary mx-1" id="btnRice" @click="clickedRice">Rice</button>

        <button
          type="button"
          class="btn btn-secondary mx-1"
          id="btnAsparagus"
          @click="clickedAsparagus"
        >Asparagus</button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

@Component
export default class sideMenu extends Vue {
  private clickedFrenchFries() {
    console.log("Clicked French Fries");
  }
  private clickedBakedPotato() {
    console.log("Clicked Baked Potato");
  }
  private clickedBroccoli() {
    console.log("Clicked Brocolli");
  }
  private clickedRice() {
    console.log("Clicked Rice");
  }
  private clickedAsparagus() {
    console.log("Clicked Asparagus");
  }
}
</script>

<style scoped>
</style>
