<template>
  <div class="dialog-with-overlay">
    <div class="overlay"></div>

    <div class="dialog">
      <div class="dialog-body">
        <span class="dialog-prompt">{{ prompt }}</span>
        <input type="text" v-model="value" v-focus v-on:keyup.enter="() => okHandler(value)" />
      </div>

      <div class="dialog-controls">
        <button class="dialog-ok-btn" @click="() => okHandler(value)">OK</button>
        <button class="dialog-cancel-btn" @click="cancelHandler">Cancel</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Prop, Component } from "vue-property-decorator";

@Component
export default class DialogBox extends Vue {
  @Prop() private readonly prompt: any;
  @Prop() private readonly okHandler: any;
  @Prop() private readonly cancelHandler: any;

  private value: number | null = null;
}
</script>

<style scoped>
.dialog-with-overlay {
  width: 100%;
  height: 100%;
  position: fixed;
  top: 0;
  left: 0;
  z-index: 10;
}

.overlay {
  width: 100%;
  height: 100%;
  position: fixed;
  top: 0;
  left: 0;
  background-color: hsla(243, 16%, 36%, 0.65);
}

.dialog {
  width: 250px;
  height: 200px;
  position: fixed;
  top: 50%;
  left: 50%;
  background-color: white;
  margin-left: -125px;
  margin-top: -100px;
}

.dialog-body {
  text-align: center;
  margin-top: 1.5rem;
}

.dialog-prompt {
  display: block;
  text-align: center;
  font-weight: bold;
  font-size: 1.5rem;
}

.dialog-controls {
  height: 25%;
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
}

button {
  height: 100%;
  width: 50%;
  background-color: #3678a8;
  color: white;
  border: 0;
}

button:hover {
  background-color: #2d648b;
}
</style>
