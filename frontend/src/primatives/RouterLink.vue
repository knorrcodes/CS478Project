<template>
  <div class="row dspl mx-1 my-1">
    <router-link class="text btn" :to="to">{{value}}</router-link>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";

@Component
export default class routerlink extends Vue {
  @Prop() private value: any;
  @Prop() private to: any;
}
</script>

<style scoped>
.text {
  background-color: #3678a8;
  color: black;
  width: 40%;
  height: 100px;
  margin-top: 5px;
  margin-bottom: 5px;
  font-size: 25px;
}

.dspl {
  display: inline;
}
</style>