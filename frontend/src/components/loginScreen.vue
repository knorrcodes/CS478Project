<template>
  <div>
    <div class="id1 container">
      <div class="row justify-content-center">
        <form action>
          <p class="text-center">Server Login</p>
          <input id="serverCodeNumber" class="form-control" v-model="serverCode" type="number" />
          <br />
          <button
            type="button"
            class="btn btn-primary"
            id="serverCodeButton"
            @click="serverCodeCheck"
          >Enter</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";

@Component
export default class loginScreen extends Vue {
  private serverCode: Number = 0;

  private serverCodeCheck() {
    if (this.serverCode == 478) {
      console.log("code is valid");
    } else {
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
  display: block;
  margin-left: auto;
  margin-right: auto;
}
</style>