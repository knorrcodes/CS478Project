<template>
  <div class="background" id="app">
    <component :is="layout">
      <router-view />
    </component>
  </div>
</template>

<script>
import { Component, Vue } from "vue-property-decorator";

const default_layout = "main";

@Component
export default class App extends Vue {
  get layout() {
    return (this.$route.meta.layout || default_layout) + "-layout";
  }
}
</script>

<style>
body {
  background-color: rgb(255, 111, 111) !important;
}
</style>
