<template>
  <div>
    <div class="id1 container">
      <div class="row justify-content-center">
        <form action>
          <p class="text-center textFormat">Server Login</p>
          <input
            id="serverCodeNumber"
            class="form-control input1 textFormat"
            v-model="serverCode"
            @keydown.enter.prevent="serverCodeCheck"
            type="number"
          />
          <div v-if="isError">
            <h4 class="errorMessage">{{error}}</h4>
          </div>
          <br />
          <button
            type="button"
            class="btn btn-primary"
            id="serverCodeButton"
            @click="serverCodeCheck"
          >Enter</button>
        </form>
        <div v-if="clicked">
          <p>{{apiData}}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import ApiClass from "../api";
import mainMenu from "./mainMenu.vue";

@Component
export default class loginScreen extends Vue {
  private serverCode: Number = 0;
  private apiData: String = "";
  clicked: boolean = false;
  private mainMenu = true;
  private isError = false;
  private error = "";

  private serverCodeCheck() {
    if (this.serverCode == 478) {
      console.log("code is valid");
      this.$router.push({ name: "mainMenu" });
    } else {
      this.isError = true;
      this.error = "code is not valid";
      console.log("code is not valid");
    }
  }
}
</script>

<style scoped>
.id1 {
  text-align: center;
  justify-content: center;
}
.input1 {
  -webkit-appearance: none;
  background: rgba(255, 255, 255, 0);
  text-align: center;
  border: none;
  display: block;
  margin-left: auto;
  margin-right: auto;
}

.textFormat {
  font-family: "Gill Sans", "Gill Sans MT", Calibri, "Trebuchet MS", sans-serif;
  font-size: 18pt;
}

input[type="number"]::-webkit-inner-spin-button,
input[type="number"]::-webkit-outer-spin-button {
  -webkit-appearance: none;
  -moz-appearance: none;
  appearance: none;
  margin: 0;
}

.errorMessage {
  color: red;
  text-transform: uppercase;
}
</style>